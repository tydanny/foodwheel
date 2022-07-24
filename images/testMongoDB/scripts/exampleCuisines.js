db = connect( 'mongodb://localhost/cuisines' );

db.cuisines.insertMany([
  {
    "Name": "North_American",
    "Dishes": ["Burgers", "Fired Chicken"]
  }, {
    "Name": "South_American",
  	"Dishes": ["Burritos", "Tacos", "Quesadillas"]
  }, {
    "Name": "Chinese",
  	"Dishes": ["Orange Chicken", "General Tso's Chicken", "Beijing Beef", "Ham Fried Rice"]
  }, {
   	"Name": "Indian",
    "Dishes": ["Chicken Tikka Masala", "Naan", "Kofta"]
  }
])
