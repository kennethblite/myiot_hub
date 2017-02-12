package sensor


import(

)

var prevSlope bool
var slope bool
type RotateBuffer struct{
	samplepoints int
	filled int
	pointer int
	length int
	sum float64
	prev_sum float64
	data []float64
}

func NewRotateBuffer(length int)*RotateBuffer{
	buffer := make([]float64,length)
	return &RotateBuffer{length:length, data:buffer}
}

func (r RotateBuffer) Average()float64{
	return r.sum/float64(r.filled)	
}

func (r RotateBuffer) SampleNumber()int{
	return r.samplepoints
}

func (r RotateBuffer) SlopeChange()bool{
	prevSlope = slope
	slope = r.sum> r.prev_sum	
	return slope != prevSlope
}


func (r *RotateBuffer) Add(f float64){
	if f == 615.25{
		return
	}	
	r.samplepoints++
	if r.filled < r.length {
		r.data[r.pointer] = f
		r.sum += f
		r.pointer++
		r.filled++ 
		return
	}
	if r.pointer == r.length {
		r.pointer = 0
	}
	r.prev_sum = r.sum
	r.sum -= r.data[r.pointer]
	r.data[r.pointer] = f
	r.sum += f
	r.pointer++
	return
}
