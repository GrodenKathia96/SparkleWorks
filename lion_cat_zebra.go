package jewelryMaking

import "fmt"

// Represents the metals and stone used in jewelry making.
type JewelryMaterial struct {
	metal string
	stone string
}

// Represents the tools used for jewelry making.
type JewelryTool struct {
	name string
	use  string
}

// Represents a single piece of jewelry being made.
type JewelryPiece struct {
	material *JewelryMaterial
	tools    []*JewelryTool
}

// Represents the jewelry-making workshop.
type JewelryWorkshop struct {
	piecesMade []*JewelryPiece
	workbench  []*JewelryTool
}

// Creates a new JewelryMaterial.
func NewJewelryMaterial(metal string, stone string) *JewelryMaterial {
	return &JewelryMaterial{
		metal: metal,
		stone: stone,
	}
}

// Creates a new JewelryTool.
func NewJewelryTool(name string, use string) *JewelryTool {
	return &JewelryTool{
		name: name,
		use:  use,
	}
}

// Creates a new JewelryPiece.
func NewJewelryPiece(material *JewelryMaterial, tools []*JewelryTool) *JewelryPiece {
	return &JewelryPiece{
		material: material,
		tools:    tools,
	}
}

// Creates a new JewelryWorkshop with the given tools.
func NewJewelryWorkshop(workbench []*JewelryTool) *JewelryWorkshop {
	return &JewelryWorkshop{
		workbench: workbench,
	}
}

// Adds a JewelryPiece to the JewelryWorkshop.
func (jw *JewelryWorkshop) AddPiece(piece *JewelryPiece) {
	jw.piecesMade = append(jw.piecesMade, piece)
}

// Prints the JewelryWorkshop's pieces.
func (jw *JewelryWorkshop) PrintPieces() {
	fmt.Println("Pieces Made:")
	for _, piece := range jw.piecesMade {
		fmt.Printf("Material: %s %s, Tools: %v\n", piece.material.metal, piece.material.stone, piece.tools)
	}
}

// Prints the JewelryWorkshop's tools.
func (jw *JewelryWorkshop) PrintTools() {
	fmt.Println("Tools:")
	for _, tool := range jw.workbench {
		fmt.Printf("Name: %s, Use: %s\n", tool.name, tool.use)
	}
}

func main() {
	// Create the tools.
	jewelerHammer := NewJewelryTool("Jeweler Hammer", "Hammering")
	polishingCloth := NewJewelryTool("Polishing Cloth", "Polishing")
	jewelersSaw := NewJewelryTool("Jeweler's Saw", "Sawing")
	jewelersPliers := NewJewelryTool("Jeweler's Pliers", "Pliers")

	// Create the workshop.
	workshop := NewJewelryWorkshop([]*JewelryTool{
		jewelerHammer,
		polishingCloth,
		jewelersSaw,
		jewelersPliers,
	})

	// Create a piece of jewelry using silver and diamond.
	silverDiamondPiece := NewJewelryPiece(
		NewJewelryMaterial("Silver", "Diamond"),
		[]*JewelryTool{
			jewelerHammer,
			polishingCloth,
			jewelersSaw,
			jewelersPliers,
		},
	)

	// Add the piece to the workshop.
	workshop.AddPiece(silverDiamondPiece)

	// Print the pieces and tools in the workshop.
	workshop.PrintPieces()
	workshop.PrintTools()
}