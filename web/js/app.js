async function searchProducts() {
    const keyword = document.getElementById('keyword').value;
    const priceMin = document.getElementById('price_min').value;
    const priceMax = document.getElementById('price_max').value;
    const category = document.getElementById('category').value;
    const sortBy = document.getElementById('sort_by').value;
    const sortOrder = document.getElementById('sort_order').value;

    const params = new URLSearchParams({
        keyword: keyword,
        price_min: priceMin,
        price_max: priceMax,
        category: category,
        sort_by: sortBy,
        sort_order: sortOrder
    });

    const response = await fetch(`http://localhost:8080/search?${params.toString()}`);
    const products = await response.json();

    const resultsDiv = document.getElementById('search_results');
    resultsDiv.innerHTML = '';

    if (products.length > 0) {
        products.forEach(product => {
            const productDiv = document.createElement('div');
            productDiv.innerHTML = `
                <h3>${product.Name}</h3>
                <p>Brand: ${product.Brand}</p>
                <p>Price: $${product.Price}</p>
                <p>Category: ${product.Category}</p>
                <p>Rating: ${product.Rating}</p>
            `;
            resultsDiv.appendChild(productDiv);
        });
    } else {
        resultsDiv.innerHTML = '<p>No products found.</p>';
    }
}

async function getSuggestions(userID) {
    const response = await fetch(`http://localhost:8080/suggestions?user_id=${userID}`);
    const suggestions = await response.json();

    const suggestionsDiv = document.getElementById('suggestions');
    suggestionsDiv.innerHTML = '';

    if (suggestions.length > 0) {
        suggestions.forEach(product => {
            const productDiv = document.createElement('div');
            productDiv.innerHTML = `
                <h3>${product.Name}</h3>
                <p>Brand: ${product.Brand}</p>
                <p>Price: $${product.Price}</p>
                <p>Category: ${product.Category}</p>
                <p>Rating: ${product.Rating}</p>
            `;
            suggestionsDiv.appendChild(productDiv);
        });
    } else {
        suggestionsDiv.innerHTML = '<p>No suggestions found.</p>';
    }
}
