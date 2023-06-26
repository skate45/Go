import (
   "sort"
)

func solution(a []int) []int {
    filtered := Filter(a, func(itemInt int) bool {
        return itemInt!=-1
    })
    
    Ordena(filtered,false)
    
    datosEnOrdenConArboles:=[]int{}
    indexEnFiltrados:=0
    
    for _,element:= range a {
        if element==-1{
            datosEnOrdenConArboles=append(datosEnOrdenConArboles,element)
        }else{
            datosEnOrdenConArboles=append(datosEnOrdenConArboles,filtered[indexEnFiltrados])
            indexEnFiltrados++
        }
    }
    
    
    return datosEnOrdenConArboles
}

func Filter(datos []int, f func(int) bool) []int {
    filtered := make([]int, 0)
    for _, v := range datos {
        if f(v) {
            filtered = append(filtered, v)
        }
    }
    return filtered
}

func Ordena(datos [] int, esOrdenAscendente bool){
    if(esOrdenAscendente){
        sort.Slice(datos, func(i, j int) bool {
            return datos[j] < datos[i]
        })
    }else{
        sort.Slice(datos, func(i, j int) bool {
            return datos[j] > datos[i]
        })
    }
}
