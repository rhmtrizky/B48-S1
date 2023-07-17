const projectData = []

function addData(program) {
    program.preventDefault()

    let projectName = document.getElementById("project-name").value
    let startDate = document.getElementById("start-date").value
    let endDate = document.getElementById("end-date").value
    let description = document.getElementById("description").value
    let image = document.getElementById("input-image").files

    // icon
    const IconNodeJS = '<i class="fa-brands fa-node-js"></i>'
    const IconReactJS = '<i class="fa-brands fa-react"></i>'
    const IconJava = '<i class="fa-brands fa-java"></i>'
    const IconGolang = '<i class="fa-brands fa-golang"></i>'

    let nodeJS = document.getElementById("nodejs").checked ? IconNodeJS : ""
    let reactJS = document.getElementById("reactjs").checked ? IconReactJS : ""
    let javaScript = document.getElementById("java").checked ? IconJava : ""
    let golang = document.getElementById("golang").checked ? IconGolang : ""

    // condition to notic users if users don't completed fill out the form
    if (projectName === "") {
        return alert("Please enter your project name")
    } else if (startDate === "") {
        return alert("Please enter your start date project")
    } else if (endDate === "") {
        return alert("Please enter your end date project  ")
    } else if (description === "") {
        return alert("Please enter your description project")
    } else if (image === "") {
        return alert("Please enter your project image")
    }

    // checkBox Checked
    let checkBox = document.querySelectorAll(".checkbox-input:checked");
    if (checkBox.length === 0) {
      alert("You must select the technologies that you use.");
      return;
    }

    // Duration 
    let dateStartInput = new Date(startDate)
    let dateEndInput = new Date(endDate)

    let timeDistance = dateEndInput - dateStartInput

    let durationSecond = Math.floor(timeDistance / 1000)
    let durationMinute = Math.floor(durationSecond / 60)
    let durationHour = Math.floor(durationMinute / 60)
    let durationDay = Math.floor(durationHour / 24)
    let durationWeek = Math.floor(durationDay / 7)
    let durationMonth = Math.floor(durationWeek / 4)
    let durationYear = Math.floor(durationMonth / 12)
    let projectDuration = ""

    // condition to manage duration of project
    if (durationDay < 7) {
        projectDuration = durationDay + " Hari"
    }
    if (durationDay >= 7) {
        projectDuration = durationWeek + " minggu " + (durationDay % 7) + " hari"
    }
    if (durationWeek >= 4) {
        projectDuration = durationMonth + " bulan " + (durationWeek % 4) + " minggu " + (durationDay % 7) + " hari"
    } 
    if (durationMonth >= 12) {
        projectDuration = durationYear + " tahun " + (durationMonth % 12) + " bulan " + (durationWeek % 4) + " minggu " + (durationDay % 7) + " hari"
    }

    // to showing image file
    image = URL.createObjectURL(image[0])

    // Push project to projectData
    let project = {
        projectName,
        description,
        startDate,
        endDate,
        projectDuration,
        durationSecond,
        nodeJS,
        reactJS,
        javaScript,
        golang,
        postAt: new Date(),
        image
    }

    projectData.push(project)
    renderProject()
    console.log(projectData)

    document.getElementById("project-name").value = "";
    document.getElementById("start-date").value = "";
    document.getElementById("end-date").value = "";
    document.getElementById("description").value = "";
    document.getElementById("nodejs").checked = false;
    document.getElementById("reactjs").checked = false;
    document.getElementById("java").checked = false;
    document.getElementById("golang").checked = false;
    document.getElementById("input-image").value = "";
}
function renderProject() {
    document.getElementById("projects").innerHTML = ''
    for (let index = 0; index < projectData.length; index++) {
        document.getElementById("projects").innerHTML += `
        <div class="base-card">
            <div class="card-content">
                <div class="card-content1">
                    <img src="${projectData[index].image}" alt="file" id="project-image">
                    <h3 id="project-title" style="padding-bottom: 3px;">${projectData[index].projectName}</h3>
                    <p class="durasi" style="font-size: 10px;">Durasi: ${projectData[index].projectDuration} </p>
                    <p style="font-size: 10px;" >Publish: ${projectData[index].endDate}</p>
                    <p class="caption" id="caption">${projectData[index].description}</p>
                </div>
                <div class="card-content2">
                    <ul>
                        <li>
                            <i>${projectData[index].nodeJS}</i>
                            <i>${projectData[index].reactJS}</i>
                            <i>${projectData[index].javaScript}</i>
                            <i>${projectData[index].golang}</i>
                        </li>
                    </ul>
                </div>
                <div class="card-content3">
                    <button type="" id="edit">edit</button>
                    <button type="" id="delete" onclick="">Delete</button>
                </div>
                <p style="font-size: 10px;">${dateProject(projectData[index].postAt)}</p>
            </div>
        </div>`
    }
}
function walkDuration(duration) {

    let timeNow = new Date()
    let timePost = duration

    let distance = timeNow - timePost

    let second = Math.floor(distance / 1000)
    let minute = Math.floor(second / 60)
    let hour = Math.floor(minute / 60)
    let day = Math.floor(hour / 24)
    let month = Math.floor(day / 30)
    let year = Math.floor(month / 12)

    //condition
    if (second >= 60) {
        return `${minute} minutes ago`
    } else if (minute >= 60) {
        return `${hour} hours ago`
    } else if (hour >= 24) {
        return `${day} days ago`
    } else if (day >= 30) {
        return `${month} months ago`
    } else if (month >= 12) {
        return `${year} years ago`
    }
    return second + " detik yang lalu"
}
setInterval(() => {
    renderProject()
}, 1000)

function dateProject(publish) {

    let dateNow = new Date();

    // let listHari = ["Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"];
    // let hari = listHari[dateNow.getDay()];

    // let listBulan = ["Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Ags", "Sep", "Okt", "Nov", "Des"];
    // let bulan = listBulan[dateNow.getMonth()];

    // let year = dateNow.getFullYear();

    let hour = dateNow.getHours();
    let minute = dateNow.getMinutes();

    if (hour < 10) {
    hour = "0" + hour;
    }
    if (minute < 10) {
    minute = "0" + minute;
    }

    return `${hour}:${minute} WIB`
    
}
