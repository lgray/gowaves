package ride

// If-else statement.
type ConditionalState struct {
	params
	prev Fsm
	/*
		Offset where true branch starts execution.
		We need this because code can look like:
		if (true) then {
			let x = throw()
			5
		} else {
			let y = throw()
			6
		}

		`X` and `y` should not be executed.

	*/
	patchTruePosition uint16
	// Same as true position.
	patchFalsePosition uint16
	// Offset where `if` code block ends.
	patchNextPosition uint16
	retAssig          uint16
	startedAt         uint16
	trueStartedAt     uint16
	falseStartedAt    uint16
	rets              []uint16

	// Clean assigments after exit.
	deferred  []Deferred
	deferreds Deferreds

	condN uniqueid
}

func (a ConditionalState) backward(v Fsm) Fsm {
	a.deferred = append(a.deferred, v.(Deferred))
	return a
}

func (a ConditionalState) Property(string) Fsm {
	panic("ConditionalState Property")
}

func (a ConditionalState) Func(name string, args []string, invoke string) Fsm {
	panic("Illegal call Func on ConditionalState")
}

func (a ConditionalState) Bytes(b []byte) Fsm {
	//a.rets = append(a.rets, a.params.b.len())
	//return constant(a, a.params, rideBytes(b))
	panic("")
}

func conditionalTransition(prev Fsm, params params, deferreds Deferreds) Fsm {
	return ConditionalState{
		prev:      prev,
		params:    params,
		startedAt: params.b.len(),
		//deferred:  make([[]Deferred, 3),
		deferreds: deferreds,
	}
}

func (a ConditionalState) Condition() Fsm {
	a.rets = append(a.rets, a.params.b.len())
	return conditionalTransition(a, a.params, a.deferreds)
}

func (a ConditionalState) TrueBranch() Fsm {
	return a
}

func (a ConditionalState) FalseBranch() Fsm {
	return a
}

func (a ConditionalState) Assigment(name string) Fsm {
	n := a.params.u.next()
	//a.assigments = append(a.assigments, n)
	a.r.set(name, n)
	return assigmentFsmTransition(a, a.params, name, n, a.deferreds)
}

func (a ConditionalState) Long(value int64) Fsm {
	a.deferred = append(a.deferred, a.constant(rideInt(value)))
	return a
}

func (a ConditionalState) Call(name string, argc uint16) Fsm {
	//a.rets = append(a.rets, a.params.b.len())
	return callTransition(a, a.params, name, argc, a.deferreds)
	//panic("")
}

func (a ConditionalState) Reference(name string) Fsm {
	//a.rets = append(a.rets, a.params.b.len())
	//return reference(a, a.params, name)
	//panic("")
	a.deferred = append(a.deferred, reference(a, a.params, name))
	return a
}

func (a ConditionalState) Boolean(value bool) Fsm {
	a.deferred = append(a.deferred, a.constant(rideBoolean(value)))
	return a
}

func (a ConditionalState) String(value string) Fsm {
	a.deferred = append(a.deferred, a.constant(rideString(value)))
	return a
}

func (a ConditionalState) Return() Fsm {
	if len(a.deferred) != 3 {
		panic("len(a.deferred) != 3")
	}
	a.condN = a.u.next()
	a.deferreds.Add(a.deferred[0], a.condN, "condition cond")
	return a.prev.backward(a) //.backward(a.startedAt, a.b.len())
}

func (a ConditionalState) Write(_ params, b []byte) {
	if len(a.deferred) != 3 {
		panic("len(a.deferred) != 3")
	}

	//condB := a.deferred[0]
	trueB := a.deferred[1]
	falsB := a.deferred[2]

	a.b.writeByte(OpRef)
	a.b.write(encode(a.condN))

	a.b.jpmIfFalse()
	a.patchTruePosition = a.b.writeStub(2)
	//a.b.write(encode(a.b.len()))
	a.patchFalsePosition = a.b.writeStub(2)
	a.patchNextPosition = a.b.writeStub(2)

	a.b.patch(a.patchTruePosition, encode(a.b.len()))
	//writeDeferred(a.params, trueB)
	//a.b.ret()
	trueB.Write(a.params, nil)
	a.b.ret()

	a.b.patch(a.patchFalsePosition, encode(a.b.len()))
	falsB.Write(a.params, nil)
	a.b.ret()

	//for _, v := range condB[1:] {
	//	v.Write(a.params)
	//}

	a.b.patch(a.patchNextPosition, encode(a.b.len()))
	//for _, v := range condB[1:] {
	//	v.Clean()
	//}
	//a.b.write(b)
	//a.b.ret()

	//writeDeferred(a.params, a.deferred)
}
func (a ConditionalState) Clean() {
	//panic("ConditionalState Clean")
}
