import { DataSource } from 'typeorm';

const AppDataSource = new DataSource({
  type: "postgres",
  host: "db",
  port: 5432,
  username: "postgres",
  password: "root",
  database: "store",

  entities: [
    "dist/**/*.entity.js"
  ],
  migrations: [
    "dist/migrations/**/*.js",
  ],
});

AppDataSource.initialize()
  .then(() => {
    console.log("Data source has been initialized!");
  })
  .catch((err) => {
    console.error("Error during data source initialization", err);
  });

export default AppDataSource;
