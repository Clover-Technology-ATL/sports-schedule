// Fetch teams from the backend API and display them
fetch('http://localhost:8080/api/teams')
  .then(response => response.json())
  .then(teams => {
    const teamList = document.getElementById('teamList');
    teams.forEach(team => {
      const listItem = document.createElement('li');
      listItem.textContent = team.name;
      teamList.appendChild(listItem);
    });
  })
  .catch(error => {
    console.error('Error fetching teams:', error);
  });

// Fetch players from the backend API and display them
fetch('http://localhost:8080/api/players')
  .then(response => response.json())
  .then(players => {
    const playerList = document.getElementById('playerList');
    players.forEach(player => {
      const listItem = document.createElement('li');
      listItem.textContent = `${player.name} (Team ID: ${player.teamId})`;
      playerList.appendChild(listItem);
    });
  })
  .catch(error => {
    console.error('Error fetching players:', error);
  });

// Fetch games from the backend API and display them
fetch('http://localhost:8080/api/games')
  .then(response => response.json())
  .then(games => {
    const gameList = document.getElementById('gameList');
    games.forEach(game => {
      const listItem = document.createElement('li');
      listItem.textContent = `Game ${game.id} - Home Team: ${game.homeTeamId}, Away Team: ${game.awayTeamId}, Date: ${game.date}, Location: ${game.location}`;
      gameList.appendChild(listItem);
    });
  })
  .catch(error => {
    console.error('Error fetching games:', error);
  });

// Fetch tournaments from the backend API and display them
fetch('http://localhost:8080/api/tournaments')
  .then(response => response.json())
  .then(tournaments => {
    const tournamentList = document.getElementById('tournamentList');
    tournaments.forEach(tournament => {
      const listItem = document.createElement('li');
      listItem.textContent = tournament.name;
      tournamentList.appendChild(listItem);
    });
  })
  .catch(error => {
    console.error('Error fetching tournaments:', error);
  });