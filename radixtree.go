package radixtree

/*
	This Radix Tree is extract from Echo framework router module.
	Shortcome: Only support string with the same prefix, due to it is design for
		URL of http request.(such as: /api, /test, /xxx)
	TODO: delete API/walk API
*/
import (
	"log"
)

type RadixTree struct {
	tree *node
}

type node struct {
	label    byte
	prefix   string
	parent   *node
	children children
	ppath    string
	pnames   []string
}

type children []*node

//NewTree returns a new RadixTree instance.
func NewRadixTree() *RadixTree {
	return &RadixTree{
		tree: &node{},
	}
}
func (t *RadixTree) Add(path string) {
	if path == "" {
		panic("path cannot be empty")
	}

	pnames := []string{}
	ppath := path
	t.insert(path, ppath, pnames)
}

func (t *RadixTree) insert(path, ppath string, pnames []string) {
	cn := t.tree
	if cn == nil {
		panic("invalid tree")
	}
	search := path

	for {
		sl := len(search)
		pl := len(cn.prefix)
		l := 0

		//LCP
		max := pl
		if sl < max {
			max = sl
		}
		for ; l < max && search[l] == cn.prefix[l]; l++ {
		}

		if l == 0 {
			//At root node
			cn.label = search[0]
			cn.prefix = search
		} else if l < pl {
			//Split node
			n := newNode(cn.prefix[l:], cn, cn.children, cn.ppath, cn.pnames)

			//Reset parent node
			cn.label = cn.prefix[0]
			cn.prefix = cn.prefix[:l]
			cn.children = nil
			cn.ppath = ""
			cn.pnames = nil

			cn.addChild(n)

			if l == sl {
				//At parent node
				cn.ppath = ppath
				cn.pnames = pnames
			} else {
				//Create child node
				n = newNode(search[l:], cn, nil, ppath, pnames)
				cn.addChild(n)
			}
		} else if l < sl {
			search = search[l:]
			c := cn.findChildWithLabel(search[0])
			if c != nil {
				//Go deeper
				cn = c
				continue
			}
			//Create child node
			n := newNode(search, cn, nil, ppath, pnames)
			cn.addChild(n)
		}
		return
	}
}

func newNode(pre string, p *node, c children, ppath string, pnames []string) *node {
	return &node{
		label:    pre[0],
		prefix:   pre,
		parent:   p,
		children: c,
		ppath:    ppath,
		pnames:   pnames,
	}
}

func (n *node) addChild(c *node) {
	n.children = append(n.children, c)
}

func (n *node) findChild(l byte) *node {
	for k, c := range n.children {
		if c.label == l {
			return n.children[k]
		}
	}
	return nil
}

func (n *node) findChildWithLabel(l byte) *node {
	for _, c := range n.children {
		if c.label == l {
			return c
		}
	}
	return nil
}

func (t *RadixTree) Find(path string) bool {
	cn := t.tree //Current node as root
	var (
		search = path
		child  *node //Child node
		//		n      int    //Param counter
		//		nn *node  //Next node
		//		ns string //Next search
	)

	//Search order static > param > any
	for {
		if search == "" {
			//goto End
			return false
		}

		pl := 0 //Prefix length
		l := 0  //LCP length

		sl := len(search)
		pl = len(cn.prefix)

		//LCP
		max := pl
		if sl < max {
			max = sl
		}

		for ; l < max && search[l] == cn.prefix[l]; l++ {
		}

		if l == pl {
			//Continue search
			search = search[l:]
		} else {

		}

		if search == "" {
			goto End
		}

		//Static node
		if child = cn.findChild(search[0]); child != nil {
			cn = child
			continue
		}

		if child = cn.findChildWithLabel(search[0]); child != nil {
			cn = child
			continue
		}

		if search != "" {
			log.Printf("Not find %s\n", search)
			return false
		}

	End:
		//		log.Printf("ppath=%s\n", cn.ppath)
		return true
	}
}

func (n *node) walkNode() {
	for _, v := range n.children {
		log.Printf("label=%d, prefix=%s, ppath=%s\n", v.label, v.prefix, v.ppath)
		v.walkNode()
	}

}
func (t *RadixTree) WalkAll() {
	log.Printf("label=%d, prefix=%s, ppath=%s\n", t.tree.label, t.tree.prefix, t.tree.ppath)
	for _, v := range t.tree.children {
		log.Printf("label=%d, prefix=%s, ppath=%s\n", v.label, v.prefix, v.ppath)
		v.walkNode()
	}
}

func (c children) Len() int {
	return len(c)
}
func remove(slice children, elem *node) children {
	if slice.Len() == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			if v.children.Len() > 0 && slice[i].ppath != "" {
				slice[i].ppath = ""
			} else {
				log.Printf("slice=%p", slice)
				slice = append(slice[:i], slice[i+1:]...)
				//return remove(slice, elem)
				log.Printf("remove ok %d, prefix=%s\n", slice.Len(), elem.prefix)
			}
			break
		}
	}

	//	for _, v := range slice {
	//		log.Printf("label=%d, prefix=%s\n", v.label, v.prefix)
	//	}
	return slice
}

func (t *RadixTree) Delete(path string) bool {
	result := t.Find(path)
	if result == false {
		return false
	}

	cn := t.tree //Current node as root
	var (
		search = path
		child  *node //Child node
		parent *node = nil
		//		n      int    //Param counter
		//		nn *node  //Next node
		//		ns string //Next search
	)

	//Search order static > param > any
	for {
		if search == "" {
			//goto End
			return false
		}

		pl := 0 //Prefix length
		l := 0  //LCP length

		sl := len(search)
		pl = len(cn.prefix)

		//LCP
		max := pl
		if sl < max {
			max = sl
		}

		for ; l < max && search[l] == cn.prefix[l]; l++ {
		}

		if l == pl {
			//Continue search
			search = search[l:]
		} else {
			//			cn = nn
			//			search = ns

		}

		if search == "" {
			goto End
		}

		//Static node
		if child = cn.findChild(search[0]); child != nil {
			parent = cn
			//			log.Printf("child=%p, cn=%p\n", child, cn)
			cn = child
			continue
		}

		if search != "" {
			//			log.Printf("Not find %s\n", search)
			return false
		}

	End:
		//		log.Printf("ppath=%s\n", cn.ppath)
		if parent != nil {
			//			log.Printf("parent=%p, parent.children=%p, cn=%p\n",
			//				parent, parent.children, cn)
			parent.children = remove(parent.children, cn)

		} else {
			log.Println("parent is nil")
			if cn.ppath == path {
				cn.ppath = ""
			}
		}

		return true
	}
}
