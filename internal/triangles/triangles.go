package triangles

type Triangles []triangle

var triangles Triangles

func (t Triangles) clear() {
	triangles = nil
}

func (t Triangles) subDivide() {
	for i := len(triangles) - 1; i >= 0; i-- {
		tt := triangles[i]

		// Remove triangle with index i
		triangles = append(triangles[:i], triangles[i+1:]...)

		// Subdivide the triangle into 4 new triangles
		t1, t2, t3, t4 := tt.subDivide()

		// Add new triangles to list
		triangles = append(triangles, t1, t2, t3, t4)
	}
}
