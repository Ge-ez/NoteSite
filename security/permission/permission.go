package permission

import (
	"strings"
)

type userpermission struct {
	roles   []string
	methods []string
}

type authority map[string]userpermission

var authorities = authority{
	"/homepage": userpermission{
		roles:   []string{"student"},
		methods: []string{"GET", "POST"},
	},
	"/posts": userpermission{
		roles:   []string{"student",},
		methods: []string{"GET", "POST"},
	},
	"/uploadnote": userpermission{
		roles:   []string{"student","teacher"},
		methods: []string{"POST"},
	},
	"/editprofile": userpermission{
		roles:   []string{"student"},
		methods: []string{"POST"},
	},
	"/myprofile": userpermission{
		roles:   []string{"student","teacher"},
		methods: []string{"GET"},
	},
	"/adminhomepage": userpermission{
		roles:   []string{"admin"},
		methods: []string{"GET"},
	},
	"/logout": userpermission{
    		roles:   []string{"student","teacher"},
    		methods: []string{"POST"},
    	},
}
