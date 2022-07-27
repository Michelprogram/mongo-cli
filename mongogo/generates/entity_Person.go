package //Your package name

type Person struct{

    
        Age int `json:"age" bson:"age"`
    
        NumeroDeTelephone string `json:"numero_de_telephone" bson:"numero_de_telephone"`
    
        Beekeeper *Beekeeper `json:"beekeeper" bson:"beekeeper"`
    

}