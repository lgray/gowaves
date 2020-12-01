package ride

import (
	"github.com/pkg/errors"
	"github.com/wavesplatform/gowaves/pkg/proto"
	"go.uber.org/zap"
)

func CallVerifier2(env RideEnvironment, tree *Tree) (RideResult, error) {
	e, err := treeVerifierEvaluator(env, tree)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call verifier")
	}
	return e.evaluate()
}

func CallVerifier3(txID string, env RideEnvironment, tree *Tree) (RideResult, error) {
	compiled, err := CompileVerifier(txID, tree)
	if err != nil {
		return nil, errors.Wrap(err, "call compile script")
	}
	return compiled.Run(env)
}

func CallVerifier(txID string, env RideEnvironment, tree *Tree) (RideResult, error) {
	r, err := CallVerifier3(txID, env, tree)
	if err != nil {
		return nil, err
	}

	//r, err := CallVerifier2(env, tree)
	//if err != nil {
	//	return nil, err
	//}
	//if !r.Eq(r2) {
	//	return nil, errors.New("R1 != R2: failed to call account script on transaction ")
	//}
	return r, nil
}

func CallFunction3(env RideEnvironment, tree *Tree, name string, args proto.Arguments) (RideResult, error) {
	if name == "" {
		name = "default"
	}
	e, err := treeFunctionEvaluator(env, tree, name, args)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to call function '%s'", name)
	}
	return e.evaluate()
}

func CallFunction(txID string, env RideEnvironment, tree *Tree, name string, args proto.Arguments) (RideResult, error) {
	rs1, err := CallFunction3(env, tree, name, args)
	if err != nil {
		return nil, errors.Wrap(err, "call function by tree")
	}
	rs2, err := CallFunction2(txID, env, tree, name, args)
	if err != nil {
		return nil, errors.Wrap(err, "call function by vm")
	}
	if !rs1.Eq(rs2) {
		zap.S().Errorf("result mismatch tree %+q  vm %+q", rs1, rs2)
		return nil, errors.New("result mismatch")
	}
	return rs2, nil
}

func CallFunction2(txID string, env RideEnvironment, tree *Tree, name string, args proto.Arguments) (RideResult, error) {
	if name == "" {
		name = "default"
	}
	f, err := CompileFunction(txID, tree, name, args)
	if err != nil {
		return nil, err
	}
	return f.Run(env)
}
