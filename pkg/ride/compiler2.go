package ride

import (
	"math"

	"github.com/pkg/errors"
)

func ccc(f Fsm, node Node) (Fsm, error) {
	switch n := node.(type) {
	case *AssignmentNode:
		fsm, err := ccc(f.Assigment(n.Name), n.Expression)
		if err != nil {
			return fsm, err
		}
		return ccc(fsm.Return(), n.Block)
	case *LongNode:
		return f.Long(n.Value), nil
	case *FunctionCallNode:
		var err error
		f = f.Call(n.Name, uint16(len(n.Arguments)))
		for i := range n.Arguments {
			f, err = ccc(f, n.Arguments[i])
			if err != nil {
				return f, err
			}
		}
		return f.Return(), nil
	case *ReferenceNode:
		return f.Reference(n.Name), nil
	case *BooleanNode:
		return f.Boolean(n.Value), nil
	case *StringNode:
		return f.String(n.Value), nil
	case *BytesNode:
		return f.Bytes(n.Value), nil
	case *ConditionalNode:
		f, err := ccc(f.Condition(), n.Condition)
		if err != nil {
			return f, err
		}
		f, err = ccc(f.TrueBranch(), n.TrueExpression)
		if err != nil {
			return f, err
		}
		f, err = ccc(f.FalseBranch(), n.FalseExpression)
		if err != nil {
			return f, err
		}
		return f.Return(), nil
	case *FunctionDeclarationNode:
		fsm, err := ccc(f.Func(n.Name, n.Arguments, n.invocationParameter), n.Body)
		if err != nil {
			return fsm, err
		}
		return ccc(fsm.Return(), n.Block)
	case *PropertyNode:
		f, err := ccc(f.Property(n.Name), n.Object)
		if err != nil {
			return f, err
		}
		return f.Return(), nil
	case nil:
		// it should be dapp
		return f, nil
	default:
		return f, errors.Errorf("unknown type %T", node)
	}
}

func CompileVerifier(txID string, tree *Tree) (*Executable, error) {
	if tree.IsDApp() {
		if tree.HasVerifier() {
			_, ok := tree.Verifier.(*FunctionDeclarationNode)
			if !ok {
				return nil, errors.New("invalid verifier declaration")
			}
			return compileFunction(txID, tree.LibVersion, append(tree.Declarations, tree.Verifier), tree.IsDApp(), tree.HasVerifier())
		}
		return nil, errors.New("no verifier declaration")
	}
	return compileFunction(txID, tree.LibVersion, []Node{tree.Verifier}, tree.IsDApp(), tree.HasVerifier())
}

func CompileDapp(txID string, tree *Tree) (*Executable, error) {
	if !tree.IsDApp() {
		return nil, errors.Errorf("unable to compile dappp")
	}
	fns := tree.Functions
	if tree.HasVerifier() {
		fns = append(fns, tree.Verifier)
	}
	return compileFunction(txID, tree.LibVersion, append(tree.Declarations, fns...), true, tree.HasVerifier())
}

func compileFunction(txID string, libVersion int, nodes []Node, isDapp bool, hasVerifier bool) (*Executable, error) {
	fCheck, err := selectFunctionChecker(libVersion)
	if err != nil {
		return nil, err
	}
	u := &uniqid{}
	b := newBuilder()
	r := newReferences(nil)
	c := newCell()
	b.writeByte(OpReturn)

	params := params{
		b:    b,
		r:    r,
		f:    fCheck,
		u:    u,
		c:    c,
		txID: txID,
	}
	for k, v := range predefinedFunctions {
		params.addPredefined(v.name, uint16(math.MaxUint16-k), uint16(math.MaxUint16-k))
	}

	f := NewMain(params)
	for _, node := range nodes {
		f, err = ccc(f, node)
		if err != nil {
			return nil, err
		}
	}
	// Just to write `OpReturn` to bytecode.
	f = f.Return()

	return f.(BuildExecutable).BuildExecutable(libVersion, isDapp, hasVerifier), nil
}

func CompileTree(tx string, tree *Tree) (*Executable, error) {
	if tree.IsDApp() {
		return CompileDapp(tx, tree)
	}
	return CompileVerifier(tx, tree)
}