package cubes


type hitung2d interface {
    Luas() float64
    Keliling() float64
}

type hitung3d interface {
    volume() float64
}

type Hitung interface {
    Hitung2d() float64
    Hitung3d() float64
}

type Cube struct {
    P float64
    L float64
    T float64
}

func (c Cube) Luas() float64 {
    return c.P*c.L
}
func (c Cube) Keliling() float64 {
    return 2*(c.P+c.T)
}
func (c Cube) Volume() float64 {
    return c.P*c.L*c.T
}