package eval

import (
	"fmt"
	"math"
	"strings"
)

// Expr 算数表达式
type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error // Check方法报告表达中的错误，并把表达式中的变量加入Vars中
}

// Var 表示一个变量，比如 x
type Var string

// lister 是一个数字常量，比如 3.141
type literal float64

// unary 表示一元操作符表达式，比如 -x
type unary struct {
	op rune // '+', '-' 中的一个
	x  Expr
}

// binary 表示二元操作符表达式，比如 x+y
type binary struct {
	op   rune // '+', '-', '*', '/' 中的一个
	x, y Expr
}

// call 表示函数调用表达式，比如 sin(x)
type call struct {
	fn   string // one of "pow", "sin", "sqrt" 中的一个
	args []Expr
}

// Env 上下文环境变量
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	// map类型的另一种声明方式
	vars := map[Var]bool{}
	// vars := make(map[Var]bool)
	if err := c.Check(vars); err != nil {
		fmt.Printf("check args error: %s\n", err)
	}
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary of op %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binay op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}
