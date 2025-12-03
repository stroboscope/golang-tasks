package reflect_spell

import "reflect"

type Spell interface {
	// название заклинания
	Name() string
	// характеристика, на которую воздействует
	Char() string
	// количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу, то заклинание применяется через него
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object interface{}) {

	v := reflect.ValueOf(object)

	if v.IsValid() || !v.IsNil() {
		// проверяем наличие у объекта метода ReceiveSpell
		// и применяем заклинание через него, если он присутствует
		_, ok := reflect.TypeOf(object).MethodByName("ReceiveSpell")
		if ok {
			a := make([]reflect.Value, 1)
			a[0] = reflect.ValueOf(spell)
			reflect.ValueOf(object).MethodByName("ReceiveSpell").Call(a)
			return
		}

		v = v.Elem()

		// меняем у объекта свойство, соответствующее заклинанию
		fn := v.FieldByName(spell.Char())
		if fn.CanSet() {
			fn.SetInt(fn.Int() + int64(spell.Value()))
		}
	}

}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	// nolint: unused
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}
