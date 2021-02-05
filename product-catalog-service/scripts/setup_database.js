/**
 * The mongo container image provides /docker-entrypoint-initdb.d/ path to deploy
 * a custom .js/.sh setup script that will be run once database initialisation. 
 * 
 * Initially those script will be executed against the test database by default or 
 * to the MONGO_INITDB_DATABASE if it was defined in the environment variables, which
 * in our case it is product-catalog-database.
 */

db.createUser(
    {
        user: "pc-admin",
        pwd: "pc-pass",
        roles: [
            {
                role: "readWrite",
                db : "product-catalog-database"
            }
        ]
    }
)

// Check if the new user can authenticate
db.auth("pc-admin", "pc-pass")

// Select the new database first
db = new Mongo().getDB("product-catalog-database");

// Then create all collections
db.createCollection("categories", { capped: false, autoIndexID: true })
db.createCollection("products", { capped: false, autoIndexID: true })
