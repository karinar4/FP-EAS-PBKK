package product

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/brand"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/category"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/image"
	"gorm.io/gorm"
)

var categories []*category.CategoryModel
var brands []*brand.BrandModel
var products []*ProductModel

func ImageSeeds(db *gorm.DB) {
	// Fetch all products from the database
	var products []Product
	err := db.Find(&products).Error
	if err != nil {
		fmt.Println("Error fetching products:", err)
		return
	}

	// Check if there are products available
	if len(products) == 0 {
		fmt.Println("No products found. Skipping image seeding.")
		return
	}

	// Define reusable image URL pairs
	imagePairs := [][]string{
		{"https://i.imgur.com/qgrKHCJ.jpeg", "https://i.imgur.com/TnwmiBp.png"},
		{"https://i.imgur.com/Z3cZQz8.jpeg", "https://i.imgur.com/vqFopnJ.png"},
		{"https://i.imgur.com/tDs0iXx.png", "https://i.imgur.com/hYWZdCw.jpeg"},
		{"https://i.imgur.com/01butBo.jpeg", "https://i.imgur.com/Ijpf8AF.jpeg"},
		{"https://i.imgur.com/Dl88mUY.jpeg", "https://i.imgur.com/u5EOZYP.png"},
	}

	// Prepare images for all products
	var images []*image.ImageModel
	for i, product := range products {
		// Determine the pair of images to use by cycling through the list
		imagePair := imagePairs[i%len(imagePairs)]

		// Add two images for the current product
		images = append(images, &image.ImageModel{
			BaseModels: common.BaseModels{ID: uuid.New()},
			ProductID:  product.ID,
			URL:        imagePair[0],
		})
		images = append(images, &image.ImageModel{
			BaseModels: common.BaseModels{ID: uuid.New()},
			ProductID:  product.ID,
			URL:        imagePair[1],
		})
	}

	// Insert all images into the database
	err = db.Create(images).Error
	if err != nil {
		fmt.Println("Error creating images:", err)
	} else {
		fmt.Println("Successfully seeded images for all products")
	}
}

func CategorySeeds(db *gorm.DB) {
	categories = []*category.CategoryModel{
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Photography",
		},
		{

			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Surveying",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Cinematography",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Mapping",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Underwater",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Industrial",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Inspection",
		},
	}

	err := db.Create(categories).Error
	if err != nil {
		fmt.Println("Error when create categories")
	} else {
		fmt.Println("Success create categories")
	}
}

func BrandSeeds(db *gorm.DB) {
	brands = []*brand.BrandModel{
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Parrot Anafi",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "DJI",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Autel",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Unitree",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "FLIR",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "QySea",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Inspired Fligth",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Flybotix",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Chasing",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "ACSL",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Freefly",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Skydio",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Elistair",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Hexaero",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "AerialX",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Percepto",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Exyn",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Velodyne",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "AscTec",
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: "Boston Dynamics",
		},
	}

	err := db.Create(brands).Error
	if err != nil {
		fmt.Println("Error when create brands")
	} else {
		fmt.Println("Success create brands")
	}
}

func ProductSeeds(db *gorm.DB) {
	products = []*ProductModel{
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Parrot Anafi Rental Kit",
			Description: "High-performance drone ideal for photography and thermal imaging.",
			Price:       200.00,
			Stock:       3,
			CategoryID:  categories[0].ID,
			BrandID:     brands[0].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Matrice 350 RTK + H20T Sensor",
			Description: "Industrial-grade drone with advanced thermal and zoom capabilities.",
			Price:       525.00,
			Stock:       3,
			CategoryID:  categories[1].ID,
			BrandID:     brands[1].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Evo 2 Pro 6k",
			Description: "Professional cinematography drone with long battery life and 6K camera.",
			Price:       40.00,
			Stock:       5,
			CategoryID:  categories[2].ID,
			BrandID:     brands[2].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Unitree Go2 Pro Quadruped Robot",
			Description: "Agile and intelligent quadruped robot suitable for research and exploration.",
			Price:       199.00,
			Stock:       10,
			CategoryID:  categories[3].ID,
			BrandID:     brands[3].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Teledyne FLIR SIRAS",
			Description: "Thermal imaging drone for industrial and inspection tasks.",
			Price:       249.00,
			Stock:       2,
			CategoryID:  categories[1].ID,
			BrandID:     brands[4].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "QySea Fifish V6 Expert M200",
			Description: "Underwater drone with advanced maneuverability and HD camera.",
			Price:       125.00,
			Stock:       3,
			CategoryID:  categories[4].ID,
			BrandID:     brands[5].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Inspired Flight Heavy Lifter",
			Description: "Heavy-duty industrial drone designed for high-capacity tasks.",
			Price:       2500.00,
			Stock:       2,
			CategoryID:  categories[5].ID,
			BrandID:     brands[6].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Flybotix ASIO Pro",
			Description: "Advanced drone for industrial inspection with superior stabilization.",
			Price:       4200.00,
			Stock:       2,
			CategoryID:  categories[6].ID,
			BrandID:     brands[7].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Chasing M2 Underwater Drone",
			Description: "Compact and efficient underwater exploration drone.",
			Price:       320.00,
			Stock:       5,
			CategoryID:  categories[4].ID,
			BrandID:     brands[8].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "ACSL SOTEN",
			Description: "Japanese industrial drone with robust design for various applications.",
			Price:       1900.00,
			Stock:       4,
			CategoryID:  categories[5].ID,
			BrandID:     brands[9].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Freefly Alta X",
			Description: "Heavy-duty cinematography drone with payload capabilities.",
			Price:       4500.00,
			Stock:       3,
			CategoryID:  categories[2].ID,
			BrandID:     brands[10].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Skydio X2",
			Description: "Autonomous drone with advanced AI for mapping and inspection.",
			Price:       2400.00,
			Stock:       6,
			CategoryID:  categories[6].ID,
			BrandID:     brands[11].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Elistair Orion 2",
			Description: "Tethered drone for industrial surveillance and inspection.",
			Price:       12000.00,
			Stock:       1,
			CategoryID:  categories[6].ID,
			BrandID:     brands[12].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Hexaero Surveyor Pro",
			Description: "Advanced drone for professional surveying tasks.",
			Price:       3600.00,
			Stock:       5,
			CategoryID:  categories[1].ID,
			BrandID:     brands[13].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "AerialX Hammerhead",
			Description: "Specialized drone for industrial payload delivery.",
			Price:       7500.00,
			Stock:       2,
			CategoryID:  categories[5].ID,
			BrandID:     brands[14].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Percepto AIM",
			Description: "Autonomous industrial inspection drone with AI capabilities.",
			Price:       8000.00,
			Stock:       3,
			CategoryID:  categories[6].ID,
			BrandID:     brands[15].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Exyn A3R",
			Description: "AI-powered mapping drone for industrial environments.",
			Price:       10000.00,
			Stock:       2,
			CategoryID:  categories[3].ID,
			BrandID:     brands[16].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Velodyne Lidar Drone",
			Description: "Drone with high-precision lidar for 3D mapping.",
			Price:       5000.00,
			Stock:       3,
			CategoryID:  categories[3].ID,
			BrandID:     brands[17].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "AscTec Falcon 8+",
			Description: "Professional drone for aerial photography and inspection.",
			Price:       2300.00,
			Stock:       4,
			CategoryID:  categories[0].ID,
			BrandID:     brands[18].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Boston Dynamics Spot",
			Description: "Agile robotic dog for industrial and research applications.",
			Price:       74000.00,
			Stock:       1,
			CategoryID:  categories[5].ID,
			BrandID:     brands[19].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "DJI Mini 4 Pro",
			Description: "Compact and lightweight drone for high-resolution photography.",
			Price:       599.00,
			Stock:       10,
			CategoryID:  categories[0].ID,
			BrandID:     brands[1].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Autel Robotics Dragonfish",
			Description: "Long-endurance VTOL drone for surveying and mapping.",
			Price:       9200.00,
			Stock:       3,
			CategoryID:  categories[3].ID,
			BrandID:     brands[2].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Unitree B2 Quadruped Robot",
			Description: "Next-gen research and mobility quadruped robot.",
			Price:       15000.00,
			Stock:       2,
			CategoryID:  categories[5].ID,
			BrandID:     brands[3].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "QySea Fifish W6",
			Description: "Industrial underwater ROV for deep-sea exploration.",
			Price:       3200.00,
			Stock:       4,
			CategoryID:  categories[4].ID,
			BrandID:     brands[5].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Parrot Anafi AI",
			Description: "AI-enhanced drone with 4G connectivity for professional inspection.",
			Price:       2500.00,
			Stock:       3,
			CategoryID:  categories[6].ID,
			BrandID:     brands[0].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Freefly Ember",
			Description: "High-speed cinema drone for elite filmmakers.",
			Price:       14500.00,
			Stock:       1,
			CategoryID:  categories[2].ID,
			BrandID:     brands[10].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Chasing Gladius Mini S",
			Description: "Compact underwater drone for consumer-level exploration.",
			Price:       850.00,
			Stock:       6,
			CategoryID:  categories[4].ID,
			BrandID:     brands[8].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Skydio 2+ Pro Kit",
			Description: "Autonomous drone with advanced obstacle avoidance.",
			Price:       1999.00,
			Stock:       7,
			CategoryID:  categories[0].ID,
			BrandID:     brands[11].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Elistair Safe-T 2 Tether",
			Description: "Tethered power system for unlimited drone endurance.",
			Price:       12000.00,
			Stock:       2,
			CategoryID:  categories[6].ID,
			BrandID:     brands[12].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Hexaero V-Tail Drone",
			Description: "Unique v-tail drone design for high-speed mapping.",
			Price:       5700.00,
			Stock:       3,
			CategoryID:  categories[3].ID,
			BrandID:     brands[13].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Exyn Pak Portable LiDAR",
			Description: "Handheld 3D mapping system for industrial use.",
			Price:       20000.00,
			Stock:       1,
			CategoryID:  categories[3].ID,
			BrandID:     brands[16].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Velodyne Ultra Puck",
			Description: "Compact, high-accuracy lidar for mapping and perception.",
			Price:       6000.00,
			Stock:       3,
			CategoryID:  categories[3].ID,
			BrandID:     brands[17].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "AscTec Firefly UAV",
			Description: "Lightweight drone optimized for academic research.",
			Price:       1200.00,
			Stock:       6,
			CategoryID:  categories[0].ID,
			BrandID:     brands[18].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Boston Dynamics Stretch",
			Description: "Multi-purpose mobile industrial robot.",
			Price:       50000.00,
			Stock:       1,
			CategoryID:  categories[5].ID,
			BrandID:     brands[19].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Percepto Sparrow",
			Description: "Autonomous drone-in-a-box for industrial monitoring.",
			Price:       15000.00,
			Stock:       2,
			CategoryID:  categories[6].ID,
			BrandID:     brands[15].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Flybotix Elios 3",
			Description: "Indoor drone for confined industrial inspections.",
			Price:       8000.00,
			Stock:       2,
			CategoryID:  categories[6].ID,
			BrandID:     brands[7].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "ACSL Mini Quad",
			Description: "Compact quadcopter for entry-level surveying.",
			Price:       2400.00,
			Stock:       8,
			CategoryID:  categories[3].ID,
			BrandID:     brands[9].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Teledyne FLIR Vue TZ20-R",
			Description: "Thermal imaging payload for search and rescue operations.",
			Price:       5500.00,
			Stock:       3,
			CategoryID:  categories[1].ID,
			BrandID:     brands[4].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Parrot Disco FPV",
			Description: "Fixed-wing drone with first-person view (FPV) capabilities.",
			Price:       500.00,
			Stock:       7,
			CategoryID:  categories[2].ID,
			BrandID:     brands[0].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Autel EVO Nano+",
			Description: "Ultra-light drone with HDR photography capabilities.",
			Price:       799.00,
			Stock:       12,
			CategoryID:  categories[0].ID,
			BrandID:     brands[2].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "DJI Mini 4 Pro",
			Description: "Compact drone with 4K camera and enhanced flight time.",
			Price:       699.00,
			Stock:       7,
			CategoryID:  categories[0].ID,
			BrandID:     brands[1].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Autel Evo Nano+",
			Description: "Portable drone with advanced imaging for travelers and hobbyists.",
			Price:       799.00,
			Stock:       6,
			CategoryID:  categories[0].ID,
			BrandID:     brands[2].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Skydio 2+ Sports Kit",
			Description: "Autonomous drone optimized for outdoor sports and tracking.",
			Price:       1099.00,
			Stock:       5,
			CategoryID:  categories[6].ID,
			BrandID:     brands[11].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Parrot Bebop 2 Power",
			Description: "Durable and lightweight drone for aerial photography.",
			Price:       399.00,
			Stock:       8,
			CategoryID:  categories[0].ID,
			BrandID:     brands[0].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Chasing F1 Fish Finder",
			Description: "Smart underwater drone for fish detection and exploration.",
			Price:       299.00,
			Stock:       10,
			CategoryID:  categories[4].ID,
			BrandID:     brands[8].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "DJI Inspire 3",
			Description: "Top-tier cinematography drone with dual operator mode.",
			Price:       12000.00,
			Stock:       2,
			CategoryID:  categories[2].ID,
			BrandID:     brands[1].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Hexaero Drone Vision",
			Description: "Advanced drone for surveying and environmental monitoring.",
			Price:       5600.00,
			Stock:       3,
			CategoryID:  categories[1].ID,
			BrandID:     brands[13].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Freefly Astro Mapping Kit",
			Description: "Drone package optimized for 3D mapping and modeling.",
			Price:       8500.00,
			Stock:       2,
			CategoryID:  categories[3].ID,
			BrandID:     brands[10].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Unitree B1 Quadruped",
			Description: "Heavy-duty quadruped robot for industrial applications.",
			Price:       29000.00,
			Stock:       1,
			CategoryID:  categories[5].ID,
			BrandID:     brands[3].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "QySea V6s Underwater Drone",
			Description: "Professional-grade underwater drone with robotic arm.",
			Price:       3800.00,
			Stock:       3,
			CategoryID:  categories[4].ID,
			BrandID:     brands[5].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "DJI Phantom 4 RTK",
			Description: "Surveying drone with real-time kinematic positioning system.",
			Price:       6500.00,
			Stock:       4,
			CategoryID:  categories[1].ID,
			BrandID:     brands[1].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Elistair Safe-T 2 Tether Station",
			Description: "Tethered system for continuous industrial drone operation.",
			Price:       15000.00,
			Stock:       2,
			CategoryID:  categories[5].ID,
			BrandID:     brands[12].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Velodyne Alpha Prime Sensor",
			Description: "Lidar sensor integrated with industrial drones for mapping.",
			Price:       7200.00,
			Stock:       1,
			CategoryID:  categories[3].ID,
			BrandID:     brands[17].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Boston Dynamics Atlas",
			Description: "Humanoid robot designed for advanced research and agility.",
			Price:       100000.00,
			Stock:       1,
			CategoryID:  categories[5].ID,
			BrandID:     brands[19].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "FLIR Vue TZ20-R",
			Description: "Dual thermal imaging sensor for drones.",
			Price:       2400.00,
			Stock:       3,
			CategoryID:  categories[5].ID,
			BrandID:     brands[4].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Inspired Flight IF800 Tomcat",
			Description: "High-capacity drone optimized for heavy payloads.",
			Price:       3500.00,
			Stock:       4,
			CategoryID:  categories[5].ID,
			BrandID:     brands[6].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "AscTec Neo",
			Description: "Lightweight drone for precise industrial inspections.",
			Price:       2700.00,
			Stock:       5,
			CategoryID:  categories[6].ID,
			BrandID:     brands[18].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Parrot Bluegrass Fields",
			Description: "Agricultural drone for crop monitoring and analysis.",
			Price:       4400.00,
			Stock:       2,
			CategoryID:  categories[6].ID,
			BrandID:     brands[0].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "ACSL Mini Drone",
			Description: "Compact Japanese drone for professional tasks.",
			Price:       1500.00,
			Stock:       6,
			CategoryID:  categories[0].ID,
			BrandID:     brands[9].ID,
		},
		{
			BaseModels:  common.BaseModels{ID: uuid.New()},
			Name:        "Exyn Drone Autonomous Kit",
			Description: "Industrial drone with autonomous mapping for complex environments.",
			Price:       12000.00,
			Stock:       2,
			CategoryID:  categories[3].ID,
			BrandID:     brands[16].ID,
		},
	}

	err := db.Create(products).Error
	if err != nil {
		fmt.Println("Error when create products")
	} else {
		fmt.Println("Success create products")
	}
}
