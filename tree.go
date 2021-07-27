//Задача: Вы разрабатываете текстовый редактор для
//программистов и хотите реализовать проверку
//корректности расстановки скобок. В коде могут
//встречаться скобки []{}(). Из них скобки [,{
//и ( считаются открывающими, а соответствующими им закрывающими скобками являются
//],} и ).
// случае, если скобки расставлены неправильно, редактор должен
//также сообщить пользователю первое место, где обнаружена ошибка.
//В первую очередь необходимо найти закрывающую скобку, для которой либо нет соответствующей открывающей (например, скобка ] в
//строке “]()”), либо же она закрывает не соответствующую ей открывающую скобку (пример: “()[}”). Если таких ошибок нет,
//необходимо найти первую открывающую скобку, для которой нет соответствующей закрывающей (пример: скобка ( в строке “{}([]”).
//Помимо скобок, исходный код может содержать символы латинского алфавита, цифры и знаки препинания

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type mass_tree struct {
	tree []tree
}
type tree struct {
	parents []int64
	childs  []int64
}

func (t *mass_tree) realise(parents, input []int64) {
	var childs []int64
	for p := range parents {
		for i := range input {
			if input[i] == parents[p] {
				childs = append(childs, int64(i))
			} else {
				continue
			}
		}
	}
	m := tree{
		parents: parents,
		childs:  childs,
	}
	t.tree = append(t.tree, m)
}

func convert(vhod string) []int64 {
	var input []int64
	for i := 0; i != len(vhod); i++ {
		if vhod[i] == '-' {
			input = append(input, -1)
			i = i + 1
		} else {
			el, err := strconv.Atoi(string(vhod[i]))
			if err != nil {
				continue
			} else {
				input = append(input, int64(el))
			}
		}
	}
	return input
}
func work(vhod3 int, vhod2 string) {
	var sl []int64
	t := new(mass_tree)
	input := make([]int64, int64(vhod3))
	vhod2, _ = in.ReadString('\n')
	input = convert(vhod2)
	for i := range input {
		if input[i] == -1 {
			sl = append(sl, int64(i))
			t.realise(sl, input)
		} else {
			continue
		}
	}
	sch := 0
	for {
		in := len(t.tree) - 1
		stop := false
		for i := range t.tree[in].childs {
			for k := range input {
				if t.tree[in].childs[i] == input[k] {
					t.realise(t.tree[in].childs, input)
					stop = true
					break
				} else {
					continue
				}
			}
			if stop == false {
				continue
			} else {
				break
			}
		}
		if sch != 0 && in == len(t.tree)-1 {
			break
		} else {
			sch = sch + 1
			in = in + 1
		}
	}
	fmt.Println((len(t.tree) + 1))
}

func main() {
	var vhodchislo, vhodstroka string
	in := bufio.NewReader(os.Stdin)
	vhodchislo, _ = in.ReadString('\n')
	vhod, _ := strconv.Atoi(vhod)
	vhodstroka, _ = in.ReadString('\n')
	if vhodchislo == 1 {
		fmtt.Println("1")
	} else if vhodchiclo == 0 {
		fmt.Println("0")
	} else {
		work(vhod, vhodstrola)
	}

}
