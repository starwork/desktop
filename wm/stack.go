package wm

import (
	"fyne.io/desktop"
)

type stack struct {
	clients      []desktop.Window
	mappingOrder []desktop.Window

	listeners []desktop.StackListener
}

func (s *stack) addToStack(win desktop.Window) {
	s.clients = append([]desktop.Window{win}, s.clients...)
	s.mappingOrder = append(s.mappingOrder, win)
}

func (s *stack) addToStackBottom(win desktop.Window) {
	s.clients = append(s.clients, win)
	s.mappingOrder = append(s.mappingOrder, win)
}

func (s *stack) removeFromStack(win desktop.Window) {
	pos := s.indexForWin(win)

	if pos == -1 {
		return
	}
	s.clients = append(s.clients[:pos], s.clients[pos+1:]...)

	pos = -1
	for i, w := range s.mappingOrder {
		if w == win {
			pos = i
		}
	}
	if pos == -1 {
		return
	}
	s.mappingOrder = append(s.mappingOrder[:pos], s.mappingOrder[pos+1:]...)
}

func (s *stack) indexForWin(win desktop.Window) int {
	pos := -1
	for i, w := range s.clients {
		if w == win {
			pos = i
		}
	}
	return pos
}

func (s *stack) getMappingOrder() []desktop.Window {
	return s.mappingOrder
}

func (s *stack) getClients(clients []desktop.Window) []desktop.Window {
	return s.clients
}

func (s *stack) AddWindow(win desktop.Window) {
	if win == nil {
		return
	}
	s.addToStack(win)

	for _, l := range s.listeners {
		l.WindowAdded(win)
	}
}

func (s *stack) RemoveWindow(win desktop.Window) {
	s.removeFromStack(win)

	if s.TopWindow() != nil {
		s.TopWindow().Focus()
	}

	for _, l := range s.listeners {
		l.WindowRemoved(win)
	}
}

func (s *stack) TopWindow() desktop.Window {
	if len(s.clients) == 0 {
		return nil
	}
	return s.clients[0]
}

func (s *stack) Windows() []desktop.Window {
	return s.clients
}

func (s *stack) RaiseToTop(win desktop.Window) {
	if win.Iconic() {
		return
	}
	if len(s.clients) > 1 {
		win.RaiseAbove(s.TopWindow())
	}

	if s.indexForWin(win) == -1 {
		return
	}
	s.removeFromStack(win)
	s.addToStack(win)
}
