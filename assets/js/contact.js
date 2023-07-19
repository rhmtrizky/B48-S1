function submitData(submit) {
    submit.preventDefault();
    let name = document.getElementById('name').value 
    let email = document.getElementById('email').value 
    let phone = document.getElementById('phone').value 
    let subject = document.getElementById('subject').value 
    let message = document.getElementById('message').value 

    let objectData = {
        name,
        email,
        phone,
        subject,
        message
    }
    console.log(objectData)

    if (name === "") {
        return alert('Name not yet filled out \nData should be complete!')
    } else if (email === "") {
        return alert('Email not yet filled out \nData should be complete!')
    } else if (phone === "") {
        return alert('Phone Number not yet filled out \nData should be complete!')
    } else if (subject === "") {
        return alert('Subject not yet selected \nData should be complete!')
    } else if (message === "") {
        return alert('Message not yet filled out \nData should be complete!')
    }


    const emailReceiver = "rahmatrizkyrifai@gmail.com"
    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${name} | ${subject}&body=Hello, I am ${name} a ${subject}, ${message}. This is my contact number ${phone}.`
    a.click()
}

