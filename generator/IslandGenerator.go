package generator

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
)

type blockPlacement struct {
	x       int32
	y       int16
	z       int32
	runtime uint32
}

type Flat struct {
	biome      uint32
	placements []blockPlacement
}

func NewIsland(biome world.Biome, _ []world.Block) Flat {
	const (
		groundY = int16(64)
		baseX   = int32(7)
		baseZ   = int32(7)
	)

	addColumn := func(x, z int32, blocks ...world.Block) []blockPlacement {
		result := make([]blockPlacement, 0, len(blocks))
		for i, b := range blocks {
			result = append(result, blockPlacement{
				x:       x,
				y:       groundY - int16(i),
				z:       z,
				runtime: world.BlockRuntimeID(b),
			})
		}
		return result
	}

	placements := make([]blockPlacement, 0, 24)

	placements = append(placements, addColumn(baseX+0, baseZ+0,
		block.Grass{},
		block.Dirt{},
		block.Dirt{},
	)...)

	placements = append(placements, addColumn(baseX+1, baseZ+0,
		block.Dirt{},
		block.Dirt{},
		block.Dirt{},
	)...)

	placements = append(placements, addColumn(baseX+2, baseZ+0,
		block.Dirt{},
		block.Dirt{},
		block.Dirt{},
	)...)

	placements = append(placements, addColumn(baseX+0, baseZ+1,
		block.Dirt{},
		block.Dirt{},
		block.Dirt{},
	)...)

	placements = append(placements, addColumn(baseX+1, baseZ+1,
		block.Dirt{},
		block.Dirt{},
		block.Dirt{},
		block.Bedrock{},
	)...)

	placements = append(placements, addColumn(baseX+2, baseZ+1,
		block.Dirt{},
		block.Dirt{},
		block.Dirt{},
	)...)

	placements = append(placements, addColumn(baseX+0, baseZ+2,
		block.Dirt{},
		block.Dirt{},
		block.Dirt{},
	)...)

	placements = append(placements, addColumn(baseX+1, baseZ+2,
		block.Dirt{},
		block.Dirt{},
		block.Dirt{},
	)...)

	return Flat{
		biome:      uint32(biome.EncodeBiome()),
		placements: placements,
	}
}

func (f Flat) GenerateChunk(pos world.ChunkPos, c *chunk.Chunk) {
	minY, maxY := int16(c.Range().Min()), int16(c.Range().Max())

	for x := uint8(0); x < 16; x++ {
		for z := uint8(0); z < 16; z++ {
			for y := minY; y <= maxY; y++ {
				c.SetBiome(x, y, z, f.biome)
			}
		}
	}

	chunkOriginX := pos[0] << 4
	chunkOriginZ := pos[1] << 4

	for _, placement := range f.placements {
		if placement.y < minY || placement.y > maxY {
			continue
		}
		if placement.x < chunkOriginX || placement.x >= chunkOriginX+16 {
			continue
		}
		if placement.z < chunkOriginZ || placement.z >= chunkOriginZ+16 {
			continue
		}

		localX := uint8(placement.x - chunkOriginX)
		localZ := uint8(placement.z - chunkOriginZ)

		c.SetBlock(localX, placement.y, localZ, 0, placement.runtime)
	}
}
