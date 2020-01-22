// Find All documents in collection "Payments"
db.payment.find()

// Choose one random document
db.payment.findOne()

// Find document using id
db.payment.find({"_id" : ObjectId("598dd7f2ebbccab1405ae8f9")})

// Show only fees array
db.payment.find({"_id" : ObjectId("598dd7f2ebbccab1405ae8f9")},{"fees":1})

// Show everything else except fees
db.payment.find({"_id" : ObjectId("598dd7f2ebbccab1405ae8f9")},{"fees":0})

// Find document where fees month is August and pretty-print it
db.payment.find({"fees.month":"August"}).pretty()

// Find All documents and sort by month name
db.payment.find().sort({"fees.month":1})

// Find All documents and sort by month name and limit the results
db.payment.find().sort({"fees.month":1}).limit(2)

// Counting the result set where street is "Anne" and year 2017
db.payment.find({"address.street": "Anne", "fees.year": 2017}).count()

// Retrieving unique values , currently total payments
db.payment.distinct("fees.total")

//Find all documents where myPersonalCode is larger than 30, 
//sort by timestamp DESC and limit 100 for not to overuse memory
db['Collection'].find({ $where: "this.myPersonalCode.length > 30" }).sort({"timestamp": -1}).limit(100);

//The order in which the database refers to documents on disk.
db.getCollection('Collection').find({}).sort( { $natural : -1 } )

//Find multiple objects
db.getCollection('COLLECTION').find({"myPet": {$in: ['Muri','Maasu','Mumm'] }}).limit(10)

/*
Working with conditional operators
*/
