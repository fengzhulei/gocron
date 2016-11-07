package mycron
import "testing"

func TestSafeSet( t *testing.T){
    set := NewSet();
    set.Add(54)
    if set.Len() !=1 {
        t.Error("add set faild")
    }
    if set.IsEmpty() {
        t.Error("add set faild")
    }
    set.Clear()
    if ! set.IsEmpty(){
        t.Error("Clear set faild")
    }
    set.Add(54)
    if !set.Has(54){
        t.Error("Func Has do not work")
    }
}