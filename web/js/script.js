document.addEventListener('click', function(event) {
    var isClickInsideSearchBox = document.getElementById('search-box').contains(event.target);
    var isClickInsideSuggestions = document.getElementById('suggestions').contains(event.target);

    if (!isClickInsideSearchBox && !isClickInsideSuggestions) {
        document.getElementById('suggestions').style.display = 'none';
    }
});

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
            suggestionElement.classList.add('suggestion-item');
            suggestionElement.innerHTML = artist.name;

            suggestionElement.addEventListener('click', function() {
                window.location.href = "/artists.html?id=" + artist.id;
            });

            suggestionsContainer.appendChild(suggestionElement);
        });
        suggestionsContainer.style.display = 'block';
    } else {
        suggestionsContainer.style.display = 'none';
    }
}