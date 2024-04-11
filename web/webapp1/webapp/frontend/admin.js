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
    card.classList.add("card", "mb-3");
    card.innerHTML = `
        <div class="card-body"> 
            <div class="card-body d-flex justify-content-center align-items-center w-50 h-50">
                <img src="${el.path}" class="card-img-top img-fluid " alt="...>
            </div>

            <div class="card-body d-flex justify-content-center align-items-center">
                <h5 class="card-title">${el.mName}</h5>
            </div>

            <div class="card-body d-flex justify-content-center align-items-center">
                <a href="C:/Users/njana/OneDrive/Desktop/web/webapp1/webapp/frontend/detailedview.html?id=${el.mid}&role=admin" class="btn btn-primary view-btn">VIEW</a>
            </div>
        </div>
    `;
    container.appendChild(card);
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