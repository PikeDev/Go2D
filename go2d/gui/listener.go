 package go2d
 
 type OnClickListener struct {
	onClick func(x, y int)
 }
 
 func (o *OnClickListener) SetOnClickListener(onClick func(int, int)) {
	o.onClick = onClick
 }