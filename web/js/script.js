function searchArtist() {
    var input = document.getElementById('search-input').value;
    if (input.length > 0) {
        fetch('/api/search/artists?query=' + encodeURIComponent(input))
            .then(response => response.json())
            .then(data => {
                var suggestions = data.filter(artist => artist.name.toLowerCase().includes(input.toLowerCase()));
                showSuggestions(suggestions);
            })
            .catch(error => console.error('Error:', error));
    } else {
        document.getElementById('suggestions').style.display = 'none';
    }
}

function showSuggestions(suggestions) {
    var suggestionsContainer = document.getElementById('suggestions');
    suggestionsContainer.innerHTML = '';
    if (suggestions.length > 0) {
        suggestions.forEach(artist => {
            var suggestionElement = document.createElement('div');
            suggestionElement.innerHTML = `<a href="/artists.html?id=${artist.id}">${artist.name}</a>`;
            suggestionsContainer.appendChild(suggestionElement);
        });
        suggestionsContainer.style.display = 'block';
    } else {
        suggestionsContainer.style.display = 'none';
    }
}