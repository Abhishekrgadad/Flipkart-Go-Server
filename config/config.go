package config

import() 


func Load() *config {
	port := os.Getenv("Port")
	if port = "" {
		port = ":8000"
	}

	return &config {
		Port: port , 
		MongoURI: os.Getenv("MONGO_URI")
		JwtSecret: os.Getenv("JWT_SECRET")
	}
}