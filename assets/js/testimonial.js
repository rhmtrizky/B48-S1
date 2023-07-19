const testiPromise = new Promise(function(testiResolve, testiReject) {
    const xmlHttp = new XMLHttpRequest();

    xmlHttp.open('GET', "https://api.npoint.io/77d1587814e7c85463ca", true)
    xmlHttp.onload = function() {
        if (xmlHttp.status == 200) {
            testiResolve(JSON.parse(xmlHttp.responseText))
        } else (
            testiReject("Load data is error")
        )
    }
    xmlHttp.onerror = function() {
        testiReject("Error")
    }
    xmlHttp.send()
})

// asycn await 
let dataOfTestimonial = []
async function getTestiData(rating) {
    try {
        const response = await testiPromise
        console.log(response)
        dataOfTestimonial = response
        allTestimonials()
    } catch {
        console.log(error)
    }
}
getTestiData()

function allTestimonials() {
    let allTestimonialHTML = ""

    dataOfTestimonial.forEach((testi) => {
        allTestimonialHTML += `
        <div class="testimonial-content">
            <div class="content-left">
                <div class="ratingStar" id="rating">
                    ${ratingStar(testi.rating)}
                    </i>
                </div>
                <i></i>
                <p id="quote-testimonial">"${testi.quote}"</p>
                <p id="user-testimonial">${testi.user}</p>
                <p id="job-testimonial">${testi.job}</p>
            </div>
            <div class="content-right">
                <img src="${testi.image}" alt="" id="image-testimonial">
            </div>
        </div>`
    })
    document.getElementById("testimonials").innerHTML = allTestimonialHTML
}
function ratingStar(rating) {
    let star = '';

    let starColor = ""
    if (rating == 1) {
        starColor = '<i class="fa-solid fa-star" style="color: red; text-shadow: 2px 2px 2px rgba(0,0,0,0.8);"></i>'
    } else if (rating == 2) {
        starColor = '<i class="fa-solid fa-star" style="color: orangered; text-shadow: 2px 2px 2px rgba(0,0,0,0.8);"></i>'
    } else if (rating == 3) {
        starColor = '<i class="fa-solid fa-star" style="color: yellowgreen; text-shadow: 2px 2px 2px rgba(0,0,0,0.8);"></i>'
    } else if (rating == 4) {
        starColor = '<i class="fa-solid fa-star" style="color: rgb(22, 177, 22); text-shadow: 2px 2px 2px rgba(0,0,0,0.8);"></i>'
    } else if (rating == 5) {
        starColor = '<i class="fa-solid fa-star" style="color: rgb(84, 84, 251); text-shadow: 2px 2px 2px rgba(0,0,0,0.8);"></i>'
    }
    for (let i = 0; i < rating; i++) {
        star += starColor;
    }
    return star;
}

function filterTestimonial(rating) {
    let filterTestimonialHTML = ""

    const filteredData = dataOfTestimonial.filter((testi) => {
        return testi.rating === rating
    })

    filteredData.forEach((testi) => {
        filterTestimonialHTML += `
        <div class="testimonial-content">
            <div class="content-left">
                <div class="ratingStar" id="rating">
                    ${ratingStar(testi.rating)}
                    </i>
                </div>
                <i></i>
                <p id="quote-testimonial">"${testi.quote}"</p>
                <p id="user-testimonial">${testi.user}</p>
                <p id="job-testimonial">${testi.job}</p>
            </div>
            <div class="content-right">
                <img src="${testi.image}" alt="" id="image-testimonial">
            </div>
        </div>`
    })
    document.getElementById("testimonials").innerHTML = filterTestimonialHTML
}