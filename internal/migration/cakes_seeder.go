package migration

var CakeSeeder = Seeder{
	Table: "cakes",
	Query: `INSERT INTO cakes (title,description,rating,image) VALUES
		('Chocolate Strawberry Cake','Yummy cake made of strawberries and chocolates',8.5,'https://preppykitchen.com/wp-content/uploads/2022/05/Chocolate-Strawberry-Cake-Feature.jpg'),
		('Cheese Cake','Authentic Italian cheese cake with cherry syrup',7.3,'https://preppykitchen.com/wp-content/uploads/2022/07/Cheesecake-feature-new.jpg'),
		('Opera Cake','Box shaped chocolate cake. Delicious.',4.3,'https://preppykitchen.com/wp-content/uploads/2022/05/Opera-Cake-Feature.jpg'),
		('Devil''s Cake','If you’re a fan of chocolate, then you will love this sinfully delicious Devil’s Food Cake. It is a decadent cake that is rich, moist, and fudgy.',6.0,'https://preppykitchen.com/wp-content/uploads/2022/03/Devils-Food-Cake-Feature1.jpg'),
		('Chocolate Chip Cake','This delicious Chocolate Chip Cake is a beautiful three-layer cake that is perfect for any occasion.',8.0,'https://preppykitchen.com/wp-content/uploads/2021/10/Chocolate-Chip-Cake-Feature.jpg'),
		('Halloweens Cake','Adorably spooky, these Halloween Cake Pops are so fun to decorate into little monsters.',3.6,'https://preppykitchen.com/wp-content/uploads/2021/10/Halloween-Cake-Pops-Feature.jpg'),
		('Pumpkin Cake Pops','These Pumpkin Cake Pops are adorable and delicious little bites of joy. Crumbled pumpkin cake mixed with cream cheese then dipped in an orange coating.',4.4,'https://preppykitchen.com/wp-content/uploads/2021/10/Pumpkin-Cake-Pops-Feature.jpg'),
		('Ice Cream Cake','You won’t believe how easy it is to make an Ice Cream Cake at home! Filled with hot fudge and crumbled Oreos, this homemade version is a show-stopper.',8.7,'https://preppykitchen.com/wp-content/uploads/2021/07/Ice-Cream-Cake-Feature.jpg'),
		('No Bake Strawberry Cheesecake','This dreamy no bake strawberry cheesecake has a light and creamy filling flavored with an easy fresh strawberry reduction all in crispy Graham cracker crust. ',2.3,'https://preppykitchen.com/wp-content/uploads/2020/05/No-Bake-Strawberry-Cheesecake-feature.jpg'),
		('Vanilla Cake','Made from scratch with only a handful of ingredients, this Vanilla Cake Recipe is the perfect dessert for any occasion.',10.0,'https://preppykitchen.com/wp-content/uploads/2021/04/Vanilla-Cake-Feature.jpg')`,
}
