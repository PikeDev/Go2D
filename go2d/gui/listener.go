package go2d
 
type OnClickListener struct {
	onClick func(x, y int)
}
 
func (o *OnClickListener) SetOnClickListener(onClick func(int, int)) {
	o.onClick = onClick
}
 
type OnKeyDownListener struct {
	onKeyDown func(button int)
}
 
func (o *OnKeyDownListener) SetOnKeyDownListener(onKeyDown func(int)) {
	o.onKeyDown = onKeyDown
}
 
type OnScrollChangeListener struct {
	onScrollChange func(scrolledX, scrolledY int)
}
 
func (o *OnScrollChangeListener) SetOnScrollChangeListener(onScrollChange func(scrolledX, scrolledY int)) {
	o.onScrollChange = onScrollChange
}

type OnValueChangeListener struct {
	onValueChange func(value int)
}
 
func (o *OnValueChangeListener) SetOnValueChangeListener(onValueChange func(value int)) {
	o.onValueChange = onValueChange
}