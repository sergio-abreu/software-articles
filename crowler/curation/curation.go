package curation

const (
	MartinFowlerBlog     = "https://martinfowler.com"
	UncleBobBlog         = "https://blog.cleancoder.com"
	KamilGrzybekBlog     = "http://www.kamilgrzybek.com"
	VladimirKhorikovBlog = "https://enterprisecraftsmanship.com"
)

const (
	MartinFowler     = "Martin Fowler"
	UncleBob         = "Uncle Bob"
	KamilGrzybek     = "Kamil Grzybek"
	VladimirKhorikov = "Vladimir Khorikov"
)

func GetCuratorName(blog string) string {
	switch blog {
	case MartinFowlerBlog:
		return MartinFowler
	case UncleBobBlog:
		return UncleBob
	case KamilGrzybekBlog:
		return KamilGrzybek
	case VladimirKhorikovBlog:
		return VladimirKhorikov
	default:
		return ""
	}
}

func GetBlog(name string) string {
	switch name {
	case MartinFowler:
		return MartinFowlerBlog
	case UncleBob:
		return UncleBobBlog
	case KamilGrzybek:
		return KamilGrzybekBlog
	case VladimirKhorikov:
		return VladimirKhorikovBlog
	default:
		return ""
	}
}
