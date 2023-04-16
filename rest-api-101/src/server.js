const express = require("express");
const cors = require("cors");
const morgan = require("morgan");
const shortid = require("shortid");
const fs = require("fs/promises");
const path = require("path");

const app = express();

app.use(cors());
app.use(morgan("dev"));
app.use(express.json());

app.use((err, req, res, next) => {
  console.error(err.stack);
  res.status(500).send("Internal Server Error");
});

/**
 * Player Microservice
 * CRUD   -         Create Read Update Delete
 * GET    - /       - find all players
 * POST   - /       - create a new player and save into db
 * GET    - /:id    - find a single player
 * PUT    - /:id    - update or create player
 * PATCH  - /:id    - update player
 * DELETE - /:id    - delete player from db
 */

const dbLocation = path.resolve("src", "data.json");

// get all file data
const getDBData = async () => {
  const data = await fs.readFile(dbLocation);
  return JSON.parse(data);
};

// update player function
const updatePlayer = async (id, req) => {
  const players = await getDBData();

  const player = players.find((item) => item.id === id);

  if (!player) return false;

  player.name = req.body.name || player.name;
  player.country = req.body.country || player.country;
  player.rank = req.body.rank || player.rank;

  await fs.writeFile(dbLocation, JSON.stringify(players));

  return player;
};

// delete player from db
app.delete("/:id", async (req, res) => {
  const id = req.params.id;
  const players = await getDBData();
  const player = players.find((item) => item.id === id);
  if (!player) return res.status(404).json({ message: "Player Not Found" });

  const NewPlayers = players.filter((item) => item.id !== id);
  await fs.writeFile(dbLocation, JSON.stringify(NewPlayers));

  res.status(204).send();
});

// update or create player
app.put("/:id", async (req, res) => {
  const id = req.params.id;
  let player = await updatePlayer(id, req);

  if (!player) {
    player = {
      ...req.body,
      id: shortid.generate(),
    }

    const players = await getDBData();
    players.push(player);
    await fs.writeFile(dbLocation, JSON.stringify(players));
  }

  res.status(200).json(player);
});

// update player
app.patch("/:id", async (req, res) => {
  const id = req.params.id;
  const player = await updatePlayer(id, req);
  if (!player) return res.status(404).json({ message: "Player Not Found" });

  res.status(200).json(player);
});

// find a single player
app.get("/:id", async (req, res) => {
  const id = req.params.id;
  const players = await getDBData();

  const player = players.find((item) => item.id === id);

  if (!player) return res.status(404).json({ message: "Player Not Found" });

  res.status(201).json(player);
});

// create a new player and save into db
app.post("/", async (req, res) => {
  const player = {
    ...req.body,
    id: shortid.generate(),
  };

  const players = await getDBData();
  players.push(player);

  await fs.writeFile(dbLocation, JSON.stringify(players));

  res.status(201).json(player);
});

// find all players
app.get("/", async (_req, res) => {
  const players = await getDBData();
  res.status(201).json(players);
});

app.get("/health", (_req, res) => {
  res.status(200).json({ status: "OK" });
});

const port = process.env.PORT || 4000;
app.listen(port, () => {
  console.log(`Server listening on port ${port}`);
  console.log(`localhost: ${port}`);
});
