package product

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/brand"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/category"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"gorm.io/gorm"
)

var categories []*category.CategoryModel
var brands []*brand.BrandModel
var products []*ProductModel

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
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name:        "Parrot Anafi Rental Kit",
			Description: "Manufactured in the US - ANAFI USA offers the same high-end security, durability, and imaging capabilities as Parrot's Short-Range Reconnaissance (SRR) drone designed for the US Army. ANAFI USA's data encryption and privacy features are compliant with the European Union's General Data Protection Regulation (GDPR), delivering best-in-class privacy and security for sensitive missions. When firefighters arrive on the scene of a fire, the most important need is to view hot spots while also being able to assess the entire visual scene. ANAFI USA's gimbal and advanced optics were designed with this challenge in mind. The 32x zoom is designed around two 21-megapixel cameras, allowing operators to see details clearly from up to 5 km (3.1 mi) away. The zoom image is coupled (blended) with images from ANAFI USA's FLIR camera. This enables operators to detect hot spots with the thermal camera, while the visual camera allows them to view people and other important details from up to 2 km (1.2 mi) away. Ensuring excellent image stabilization on 32x zoom images with a light drone (500 g/1.1 lbs) is a particularly delicate technological achievement. Parrot accomplished this feat by coupling the stabilization of the gimbal with full 3-axis digital stabilization via image processing. ANAFI USA works indoors without GPS, allowing operators to take off inside a house, pilot the drone through a window to fly outside, and then come back. ANAFI USA can also launch from the palm of the hand like a paper plane, further expanding its ease of use. With no built-in limitations for no-fly zones, ANAFI USA gives first responders the freedom to fly responsibly at a moment's notice and wherever their missions may take them. ANAFI USA is the quietest drone in its class, with a sound level of just 79 dB when it is at 50cm off the ground. It uses a standard USB-C type charger for hassle-free convenience. Weighing in at just 500 g (1.1lb) the compact ANAFI USA folds easily for maximum portability. Despite its compact design, ANAFI USA boasts a 32-minute flight time - also best in class for a drone of its size. Photos and videos are encrypted on its SD card, using an AES-XTS algorithm with a 512-bit key length. ANAFI USA's Secure Digital (SD) card encryption feature ensures that saved data cannot be read if the drone or SD card is lost or stolen. ANAFI USA also includes a secure WPA2 Wi-Fi connection. WPA2 provides authentication and encryption of the link between the remote controller and ANAFI USA. ANAFI USA's advanced flight features are designed to meet the unique needs of enterprise drone users. Powered by Parrot's acclaimed FreeFlight 6 piloting software, pilots have multiple easy-to-use flight options. ANAFI USA's flight routes can be set by coordinates in assisted framing or executed autonomously. Features: - A CMOS 1/2.4'' sensor and EO tele-camera for x32 stabilized zoom - FLIR Boson® 320 longwave infrared thermal camera - Up to 4K HDR in visible spectrum and 1280x720p in thermal - WPA2 encryption and AES 512-bit secure media IN THE BOX: - 1 ANAFI USA drone - 3 smart batteries - 1 Skycontroller 3 - 1 tablet holder - 1 multi-port fast USB charger - 1 Power cord - 1 additional set of propeller blades - 4 USB-A/USB-C cables - 1 hard case",
			Price:       200.00,
			Stock:       3,
			CategoryID:  categories[0].ID,
			BrandID:     brands[0].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name:		 "Matrice 350 RTK + H20T Sensor",
			Description: "The Matrice 350 RTK takes inspiration from modern aviation systems. Offering up to 55 minutes of flight time, advanced AI capabilities, 6 Directional Sensing & Positioning, and more, the M350 RTK sets a whole new standard by combining intelligence with high-performance and unrivaled reliability. Combine the flight time with the clarity and zoom of the H20T for a superior experience with the added benefit of a 640x512px Radiometric Thermal sensor. IN THE BOX: - 1 x Matrice 350 Aircraft Body - 1 x DJI RC Plus Controller - 1 x USB Charger (wall adapter) - 1 x USB A - USB A Cable - 1 x USB C - USB A Cable - 1 x USB C - USB C Cable - 2 x TB65 batteries - 1 x BS65 Battery Charging Station - 1 x WB37 Intelligent Battery - 2 x 2110 Propeller (CW) - 2 x 2110 Propeller (CCW) - 2 x Landing Gear - 1 x Carrying Case - 1 x Manual - 1 x H20T Sensor",
			Price:		 525.00,
			Stock:		 3,
			CategoryID:  categories[1].ID,
			BrandID: 	 brands[1].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "Evo 2 Pro 6k",
			Description: "Evolve with the EVO. After countless hours of planning and testing, the EVO 2 is here - the most powerful and advanced foldable drone on the planet. Key Features: - Longest Battery Life: Up to 40 minutes of flight time. - Exceptional Performance: - Top speed of 45 mph. - Up to 9 km range. - Advanced Obstacle Avoidance: - 12 computer vision sensors, 2 sonar sensors, and 2 LED landing lights. - Omni-directional obstacle avoidance for unmatched safety. - Intelligent Features: - Failsafe alerts for low battery and return-to-home. - Autel's Dynamic Track 2.0™ for autonomous subject tracking around obstacles. - Real-time AI-powered smart flight paths. - High-Quality Camera: - EVO 2 Pro: 1'' 6K camera with adjustable aperture (f/2.8 - f/11). - Ample Storage: - 8GB internal storage. - SD card support up to 256GB. - Dual-Core Object Detection: Recognizes up to 64 subjects, including people, vehicles, and animals. What's in the Box: - Aircraft and Gimbal Cover x 1 - Battery x 2 - Remote Control x 1 - Charger x 1 - 12'' Cable USB Micro A to USB Type C x 1 - Propellers (pair) x 4 - EVO II Hard Rugged Case x 1 - 12'' Cable USB Micro A to USB Micro B x 1 - 12V Car Charger x 1 - Quick Guide",
			Price: 		 40.00,
			Stock:		 5,
			CategoryID:  categories[2].ID,
			BrandID:	 brands[2].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "Unitree Go2 Pro Quadruped Robot",
			Price: 		 199.00,
			Stock:	 	 10,
			CategoryID:  categories[3].ID,
			BrandID: 	 brands[3].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "Teledyne FLIR SIRAS",
			Price: 		 249.00,
			Stock:	 	 2,
			CategoryID:  categories[1].ID,
			BrandID: 	 brands[4].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "QySea Fifish V6 Expert M200",
			Price: 		 125.00,
			Stock:	 	 3,
			CategoryID:  categories[4].ID,
			BrandID: 	 brands[5].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "Inspired Flight",
			Price: 		 2500.00,
			Stock:	 	 2,
			CategoryID:  categories[5].ID,
			BrandID: 	 brands[6].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "Matrice 300 RTK",
			Price: 		 400.00,
			Stock:	 	 3,
			CategoryID:  categories[3].ID,
			BrandID: 	 brands[1].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "Flybotix ASIO Pro",
			Price: 		 4200.00,
			Stock:	 	 2,
			CategoryID:  categories[6].ID,
			BrandID: 	 brands[7].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "Flybotix ASIO X",
			Price: 		 900.00,
			Stock:	 	 2,
			CategoryID:  categories[6].ID,
			BrandID: 	 brands[7].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "DJI - Neo Combo",
			Price: 		 10.00,
			Stock:	 	 6,
			CategoryID:  categories[0].ID,
			BrandID: 	 brands[1].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "DJI Mavic 3 Multispectral Bundle Kit",
			Price: 		 199.00,
			Stock:	 	 10,
			CategoryID:  categories[2].ID,
			BrandID: 	 brands[1].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "DJI - Avata 2 Drone Package",
			Price: 		 65.00,
			Stock:	 	 4,
			CategoryID:  categories[0].ID,
			BrandID: 	 brands[1].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "Chasing M2 Pro Max ROV",
			Price: 		 200.00,
			Stock:	 	 2,
			CategoryID:  categories[4].ID,
			BrandID: 	 brands[8].ID,
		},
		{
			BaseModels: common.BaseModels{
				ID: uuid.New(),
			},
			Name: 		 "Soten",
			Price: 		 299.00,
			Stock:	 	 2,
			CategoryID:  categories[0].ID,
			BrandID: 	 brands[9].ID,
		},
	}

	err := db.Create(products).Error
	if err != nil {
		fmt.Println("Error when create products")
	} else {
		fmt.Println("Success create products")
	}
}
