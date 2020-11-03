//Removing data

//Remove all documents where first name is larger than 30 symbols
db['Collection'].remove( { $where: "this.FirstName.length > 30" } )

// Remove all documents where exists property "myName" : "Kaido"
db.getCollection('Collection').remove( { "myName" : "Kaido" } )
