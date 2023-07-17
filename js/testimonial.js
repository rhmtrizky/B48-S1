// OOP PROJECT

// class testimonial {
//     #quote = "";
//     #job = "";
//     #image = "";

//     constructor(quote, job, image) {
//         this.#quote = quote
//         this.#job = job
//         this.#image = image
//     }
//     get quote() {
//         return this.#quote
//     }
//     get job() {
//         return this.#job
//     }
//     get image() {
//         return this.#image
//     }
//     get user() {
//         throw new Error("You must fill out the user name at the testimonials")
//     }
//     get testimonialHTML() {
//         return `
//         <div class="testimonial-content">
//             <div class="content-left">
//                 <div class="rating">
//                     <img src="img/star-solid.svg" alt="">
//                     <img src="img/star-solid.svg" alt="">
//                     <img src="img/star-solid.svg" alt="">
//                     <img src="img/star-solid.svg" alt="">
//                     <img src="img/star-half-solid.svg" alt="">
//                 </div>
//                 <i></i>
//                 <p id="quote-testimonial">${this.quote}</p>
//                 <p id="user-testimonial">${this.user}</p>
//                 <p id="job-testimonial">${this.job}</p>
//             </div>
//             <div class="content-right">
//                 <img src="${this.image}" alt="" id="image-testimonial">
//             </div>
//         </div>`
//     }
// }

// class UserTestimonial extends testimonial {
//     #user =""

//     constructor(user, quote, job, image) {
//         super(quote, job, image)
//         this.#user = user
//     }
//     get user() {
//         return "user : " + this.#user
//     }
// }

// const testimonial1 = new UserTestimonial("Surya Elidanto", '"GG gaming"', "Frontend-Developer", "https://images.unsplash.com/photo-1580477667995-2b94f01c9516?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80")

// const testimonial2 = new UserTestimonial("Rahmat Rziky", '"Keren kamu bang"', "Backend-Developer",  "https://images.unsplash.com/photo-1541562232579-512a21360020?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// // const testimonial2 = new UserTestimonial("Guswandi", '"Keren kamu bang"', "https://images.unsplash.com/photo-1541562232579-512a21360020?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// // const testimonial3 = new CompanyTestimonial("Dumbways", "Apasih ga jelas", "https://images.unsplash.com/photo-1578632767115-351597cf2477?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// // // const testimonial4 = new Testimonial("Apasih ga jelas", "https://images.unsplash.com/photo-1578632767115-351597cf2477?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80")

// // console.log(testimonial1)

// let testimonialData = [testimonial1,testimonial2]

// let testimonialHTML = ""

// for (let i = 0; i < testimonialData.length; i++) {
//     testimonialHTML += testimonialData[i].testimonialHTML
// }

// document.getElementById("testimonials").innerHTML = testimonialHTML

// HOF AND CALLBACK
const dataOfTestimonial = [
    {
        user: "Ahmad Fauzan",
        quote: "Keren bang",
        rating: 4,
        job: "Frontend-Developer",
        image: "https://images.unsplash.com/photo-1566492031773-4f4e44671857?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80",
    },
    {
        user: "Akbar Husein",
        quote: "kelassss bang",
        rating: 5,
        job: "Web-Developer",
        image: "https://images.unsplash.com/photo-1564564321837-a57b7070ac4f?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=876&q=80"
    },
    {
        user: "Umar Zainudin",
        quote: "Biasa aja",
        rating: 3,
        job: "FullStack-Developer",
        image: "https://images.unsplash.com/photo-1539571696357-5a69c17a67c6?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80"
    },
    {
        user: "Susi Susanti",
        quote: "Jelek",
        rating: 1,
        job: "Backend-Developer",
        image: "https://images.unsplash.com/photo-1567532939604-b6b5b0db2604?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80"
    },
    {
        user: "Axel Faraday",
        quote: "Lumayan keren bang",
        rating: 3,
        job: "Frontend-Developer",
        image: "https://images.unsplash.com/photo-1590086782957-93c06ef21604?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80"
    },
    {
        user: "Farhan Yusuf",
        quote: "agak bad bang, semangatt",
        rating: 2,
        job: "UI/UX-Designer",
        image: "https://images.unsplash.com/photo-1615109398623-88346a601842?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80"
    }
];

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
allTestimonials()

function filterTestimonial(rating) {
    let filterTestimonialHTML = ""

    const filteredData = dataOfTestimonial.filter((testi) => {
        return testi.rating === rating
    })

    filteredData.forEach((testi) => {
        filterTestimonialHTML += `<div class="testimonial-content">
        <div class="content-left">
            <div class="ratingStar" id="rating">
            ${ratingStar(testi.rating)}
            </div>
            <i></i>
            <p id="quote-testimonial">"${testi.quote}"</p>
            <p id="user-testimonial">${testi.user}</p>
            <p id="job-testimonial">${testi.job}</p>
        </div>
        <div class="content-right">
            <img src="${testi.image}" alt="" id="image-testimonial">
        </div>
    </div>
    `
    })

    document.getElementById("testimonials").innerHTML = filterTestimonialHTML
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
        starColor = '<i class="fa-solid fa-star" style="color:  rgb(22, 177, 22); text-shadow: 2px 2px 2px rgba(0,0,0,0.8);"></i>'
    } else if (rating == 5) {
        starColor = '<i class="fa-solid fa-star" style="color: rgb(84, 84, 251); text-shadow: 2px 2px 2px rgba(0,0,0,0.8);"></i>'
    }
    for (let i = 0; i < rating; i++) {
        star += starColor;
    }
    return star;
}