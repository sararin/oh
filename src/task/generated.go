// Released under an MIT-style license. See LICENSE.

// Generated by generate.oh

package task

import (
	. "github.com/michaelmacinnis/oh/src/cell"
	"strings"
	"unicode"
)

func bindArithmetic(s *Scope) {

	s.DefineMethod("add", func(t *Task, args Cell) bool {
		acc := Car(args).(Number)

		for Cdr(args) != Null {
			args = Cdr(args)
			acc = acc.Add(Car(args))
		}

		return t.Return(acc)
	})

	s.DefineMethod("sub", func(t *Task, args Cell) bool {
		acc := Car(args).(Number)

		for Cdr(args) != Null {
			args = Cdr(args)
			acc = acc.Subtract(Car(args))
		}

		return t.Return(acc)
	})

	s.DefineMethod("div", func(t *Task, args Cell) bool {
		acc := Car(args).(Number)

		for Cdr(args) != Null {
			args = Cdr(args)
			acc = acc.Divide(Car(args))
		}

		return t.Return(acc)
	})

	s.DefineMethod("mod", func(t *Task, args Cell) bool {
		acc := Car(args).(Number)

		for Cdr(args) != Null {
			args = Cdr(args)
			acc = acc.Modulo(Car(args))
		}

		return t.Return(acc)
	})

	s.DefineMethod("mul", func(t *Task, args Cell) bool {
		acc := Car(args).(Number)

		for Cdr(args) != Null {
			args = Cdr(args)
			acc = acc.Multiply(Car(args))
		}

		return t.Return(acc)
	})
}

func bindGenerators(s *Scope) {

	s.DefineMethod("boolean", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(Car(args).Bool()))
	})

	s.DefineMethod("float", func(t *Task, args Cell) bool {
		return t.Return(NewFloat(Car(args).(Atom).Float()))
	})

	s.DefineMethod("integer", func(t *Task, args Cell) bool {
		return t.Return(NewInteger(Car(args).(Atom).Int()))
	})

	s.DefineMethod("pipe", func(t *Task, args Cell) bool {
		return t.Return(NewPipe(t.Lexical, nil, nil))
	})

	s.DefineMethod("rational", func(t *Task, args Cell) bool {
		return t.Return(NewRational(Car(args).(Atom).Rat()))
	})

	s.DefineMethod("status", func(t *Task, args Cell) bool {
		return t.Return(NewStatus(Car(args).(Atom).Status()))
	})

	s.DefineMethod("string", func(t *Task, args Cell) bool {
		return t.Return(NewRawString(t, Car(args).String()))
	})

	s.DefineMethod("symbol", func(t *Task, args Cell) bool {
		return t.Return(NewSymbol(raw(Car(args))))
	})
}

func bindPredicates(s *Scope) {

	s.DefineMethod("is-atom", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsAtom(Car(args))))
	})

	s.DefineMethod("is-boolean", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsBoolean(Car(args))))
	})

	s.DefineMethod("is-builtin", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsBuiltin(Car(args))))
	})

	s.DefineMethod("is-channel", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsChannel(Car(args))))
	})

	s.DefineMethod("is-cons", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsCons(Car(args))))
	})

	s.DefineMethod("is-continuation", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsContinuation(Car(args))))
	})

	s.DefineMethod("is-float", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsFloat(Car(args))))
	})

	s.DefineMethod("is-integer", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsInteger(Car(args))))
	})

	s.DefineMethod("is-method", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsMethod(Car(args))))
	})

	s.DefineMethod("is-null", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsNull(Car(args))))
	})

	s.DefineMethod("is-number", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsNumber(Car(args))))
	})

	s.DefineMethod("is-object", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsContext(Car(args))))
	})

	s.DefineMethod("is-pipe", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsPipe(Car(args))))
	})

	s.DefineMethod("is-rational", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsRational(Car(args))))
	})

	s.DefineMethod("is-status", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsStatus(Car(args))))
	})

	s.DefineMethod("is-string", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsString(Car(args))))
	})

	s.DefineMethod("is-symbol", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsSymbol(Car(args))))
	})

	s.DefineMethod("is-syntax", func(t *Task, args Cell) bool {
		return t.Return(NewBoolean(IsSyntax(Car(args))))
	})
}

func bindRelational(s *Scope) {

	s.DefineMethod("eq", func(t *Task, args Cell) bool {
		prev := Car(args)

		for Cdr(args) != Null {
			args = Cdr(args)
			curr := Car(args)

			if !prev.Equal(curr) {
				return t.Return(False)
			}

			prev = curr
		}

		return t.Return(True)
	})

	s.DefineMethod("ge", func(t *Task, args Cell) bool {
		prev := Car(args).(Number)

		for Cdr(args) != Null {
			args = Cdr(args)
			curr := Car(args).(Number)

			if prev.Less(curr) {
				return t.Return(False)
			}

			prev = curr
		}

		return t.Return(True)
	})

	s.DefineMethod("gt", func(t *Task, args Cell) bool {
		prev := Car(args).(Number)

		for Cdr(args) != Null {
			args = Cdr(args)
			curr := Car(args).(Number)

			if !prev.Greater(curr) {
				return t.Return(False)
			}

			prev = curr
		}

		return t.Return(True)
	})

	s.DefineMethod("is", func(t *Task, args Cell) bool {
		prev := Car(args)

		for Cdr(args) != Null {
			args = Cdr(args)
			curr := Car(args)

			if prev != curr {
				return t.Return(False)
			}

			prev = curr
		}

		return t.Return(True)
	})

	s.DefineMethod("le", func(t *Task, args Cell) bool {
		prev := Car(args).(Number)

		for Cdr(args) != Null {
			args = Cdr(args)
			curr := Car(args).(Number)

			if prev.Greater(curr) {
				return t.Return(False)
			}

			prev = curr
		}

		return t.Return(True)
	})

	s.DefineMethod("lt", func(t *Task, args Cell) bool {
		prev := Car(args).(Number)

		for Cdr(args) != Null {
			args = Cdr(args)
			curr := Car(args).(Number)

			if !prev.Less(curr) {
				return t.Return(False)
			}

			prev = curr
		}

		return t.Return(True)
	})
}

func bindTheRest(s *Scope) {

	s.DefineSyntax("builtin", func(t *Task, args Cell) bool {
		return t.Closure(NewBuiltin)
	})

	s.DefineSyntax("define", func(t *Task, args Cell) bool {
		return t.LexicalVar(psExecDefine)
	})

	s.DefineSyntax("dynamic", func(t *Task, args Cell) bool {
		return t.DynamicVar(psExecDynamic)
	})

	s.DefineSyntax("method", func(t *Task, args Cell) bool {
		return t.Closure(NewMethod)
	})

	s.DefineSyntax("setenv", func(t *Task, args Cell) bool {
		return t.DynamicVar(psExecSetenv)
	})

	s.DefineSyntax("syntax", func(t *Task, args Cell) bool {
		return t.Closure(NewSyntax)
	})

	s.PublicSyntax("public", func(t *Task, args Cell) bool {
		return t.LexicalVar(psExecPublic)
	})

	s.DefineMethod("car", func(t *Task, args Cell) bool {
		return t.Return(Caar(args))
	})

	s.DefineMethod("cdr", func(t *Task, args Cell) bool {
		return t.Return(Cdar(args))
	})

	s.DefineMethod("cons", func(t *Task, args Cell) bool {
		return t.Return(Cons(Car(args), Cadr(args)))
	})

	s.DefineMethod("list", func(t *Task, args Cell) bool {
		return t.Return(args)
	})

	s.DefineMethod("reverse", func(t *Task, args Cell) bool {
		return t.Return(Reverse(Car(args)))
	})
}

func bindStringPredicates(e *Env) {

	e.Method("is-control", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsControl(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-digit", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsDigit(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-graphic", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsGraphic(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-letter", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsLetter(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-lower", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsLower(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-mark", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsMark(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-print", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsPrint(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-punct", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsPunct(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-space", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsSpace(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-symbol", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsSymbol(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-title", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsTitle(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("is-upper", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		for _, c := range s {
			if !unicode.IsUpper(c) {
				return t.Return(False)
			}
		}

		return t.Return(True)
	})

	e.Method("to-lower", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		return t.Return(NewRawString(t, strings.ToLower(s)))
	})

	e.Method("to-title", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		return t.Return(NewRawString(t, strings.ToTitle(s)))
	})

	e.Method("to-upper", func(t *Task, args Cell) bool {
		s := raw(toString(Car(t.Scratch).(Binding).Self()))

		return t.Return(NewRawString(t, strings.ToUpper(s)))
	})
}
