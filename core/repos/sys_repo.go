package repos

import (
	"github.com/ixre/go2o/core/domain/interface/sys"
	impl "github.com/ixre/go2o/core/domain/sys"
	"github.com/ixre/go2o/core/infrastructure/fw"
)

var _ sys.ISystemRepo = new(systemRepoImpl)

type systemRepoImpl struct {
	fw.ORM
	areaRepo fw.Repository[sys.Region]
}

func NewSystemRepo(o fw.ORM) sys.ISystemRepo {
	return &systemRepoImpl{
		ORM: o,
	}
}

// AreaRepo implements sys.ISystemRepo.
func (s *systemRepoImpl) Region() fw.Repository[sys.Region] {
	if s.areaRepo == nil {
		s.areaRepo = newAreaRepository(s.ORM)
	}
	return s.areaRepo
}

// GetAllCities implements sys.ISystemRepo.
func (s *systemRepoImpl) GetAllCities() []*sys.Region {
	panic("unimplemented")
}

// GetRegionList implements sys.ISystemRepo.
func (s *systemRepoImpl) GetRegionList(parentId int) []*sys.Region {
	panic("unimplemented")
}

// GetSystemAggregateRoot implements sys.ISystemRepo.
func (s *systemRepoImpl) GetSystemAggregateRoot() sys.ISystemAggregateRoot {
	return impl.NewSystemAggregateRoot(s)
}

type areaRepository struct {
	fw.BaseRepository[sys.Region]
}

func newAreaRepository(o fw.ORM) fw.Repository[sys.Region] {
	s := &areaRepository{}
	s.ORM = o
	return s
}
