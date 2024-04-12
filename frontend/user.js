function loadData() {
    const container = document.getElementById("pCard");
    container.innerHTML = "";

    fetch("http://localhost:9090/viewAll")
        .then(response => {
            if (!response.ok){
                throw new Error('Failed to fetch data');
            }
            return response.json()
        })
        .then(data => {
            data.forEach(element => {
                console.log("Creating card for element:", element);
                createCard(element, container);
            });
        })
        .catch(error => console.error('Error fetching product data:', error));
}

function createCard(el, container) {
    const card = document.createElement("div");
    card.classList.add("card", "text-center", "mt-5","mx-2", "shadow", "p-3", "rounded", "col-md-2","d-flex", "flex-column","border-3", "border-dark")
    

    card.innerHTML = `
    <img src="${el.path}" class="card-img-top img-fluid" alt="Phone Image"> 
    <div class="card-body mt-auto flex-grow-1 row align-items-end">
        <a href="detailedview.html?id=${el.mid}&role=admin" class="btn btn-dark btn-sm mt-2">VIEW</a> <!-- Adjusted button size and margin -->
    </div>
    <div card-body mt-auto flex-grow-1 row align-items-end>
    <h5 class="card-title font-weight-bold text-dark">${el.mName}</h5>
    </div>
    `;
    container.appendChild(card)
}

loadData();

function search(){
    let searchInput = document.getElementById("search-input")
    const value = searchInput.value
    const keyObj = {
        key: value
    };
   
    const parsedBody = JSON.stringify(keyObj)
    console.log(parsedBody)
    fetch("http://localhost:9090/search",{
        method:"POST",
        body : parsedBody
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        return response.json();
    })
    .then(data => {
        if (data == null){
            alert('no such product')
        }
        const container = document.getElementById("pCard");
        container.innerHTML = "";
        console.log(data);
        data.forEach(element => {
                console.log("Creating card for element:", element);
                createCard(element, container);
            });
        
        data=null
        
    })
    .catch(error => {
        console.error("Error during get", error);
    });
}