package intro

type CubeService struct {
	repo CubeRepository
}

func (c *CubeService) GetCubeVolume() (float64, error) {
	return c.repo.Volume()
}

func (c *CubeService) GetCubeArea() (float64, error) {
	return c.repo.Area()
}

func NewCubeService(repo CubeRepository) CubeService {
	return CubeService{repo: repo}
}
