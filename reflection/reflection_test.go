package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with tow string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Mark", 31},
			[]string{"Mark"},
		},
		{
			"Nested fields",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{
				"Mark",
				struct {
					Age  int
					City string
				}{31, "Denver"},
			},
			[]string{"Mark", "Denver"},
		},
		{
			"Pointers to things",
			&Person{
				"Mark",
				Profile{ 31, "Denver"},
			},
			[]string{"Mark", "Denver"},
		},
		{
			"Slices",
			[]Profile{
				{31, "Denver"},
				{33, "Boulder"},
			},
			[]string{"Denver", "Boulder"},
		},
		{
			"Arrays",
			[2]Profile {
				{31, "Denver"},
				{33, "Boulder"},
			},
			[]string{"Denver", "Boulder"},
		},
		{
			"Maps",
			map[string]string {
				"Foo": "Bar",
				"Baz": "Box",
			},
			[]string{"Bar", "Box"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	t.Run("with Maps", func(t *testing.T) {
		aMap := map[string]string {
				"Foo": "Bar",
				"Baz": "Box",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Box")
	})

}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didnt", haystack, needle)
	}
}
