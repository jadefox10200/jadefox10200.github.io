<?php
    $name = $_POST['name'];
    $email = $_POST['email'];
	$message = $_POST['message'];
    $subject = $_POST['subject'];
    $from = 'From: TangledDemo';
    $to = 'jadefox10200@gmail.com';
    //$human = $_POST['human'];
    $body = "From: $name\n E-Mail: $email\n Message:\n $message";

    if ($_POST['submit']) {
        if (mail ($to, $subject, $body, $from)) {
	    echo '<p>Your message has been sent!</p>';
	} else {
	    echo '<p>Something went wrong, go back and try again!</p>';
	}
?>
