// Insert document into collection
db.cities.insert(
	[
		{"townName":"Tartu","currentPopulation":96857},
		{"townName":"Tallinn","currentPopulation":447414},
		
	]
)

// Insert some subset from Parent-Collection to Child-Collection (WHERE sentTimestamp > 1598446301802)
db.getCollection('Parent-Collection').find({"sentTimestamp":{$gt: 1598446301802}}).forEach(function(doc){
   db.getCollection('Child-Collection').insert(doc);
});
