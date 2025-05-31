package internal

import (
	"fmt"
	"testing"
)

var document = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta name="description" content="This is a brief description of the webpage content, summarizing what visitors can expect to find." />
    <title>Sample Large HTML Document with External Links</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            margin: 20px;
            max-width: 900px;
            margin-left: auto;
            margin-right: auto;
            color: #333;
        }
        header, footer {
            background-color: #f4f4f4;
            padding: 10px 20px;
            border-radius: 8px;
            margin-bottom: 20px;
        }
        nav a {
            margin-right: 15px;
            text-decoration: none;
            color: #007BFF;
        }
        nav a:hover {
            text-decoration: underline;
        }
        section {
            margin-bottom: 40px;
        }
        img {
            max-width: 100%;
            height: auto;
            border-radius: 8px;
        }
        blockquote {
            font-style: italic;
            color: #666;
            margin: 20px 40px;
            border-left: 4px solid #ccc;
            padding-left: 15px;
        }
    </style>
</head>
<body>

<header>
    <h1>Welcome to My Large HTML Document</h1>
    <nav>
        <a href="#introduction">Introduction</a>
        <a href="#history">History</a>
        <a href="#features">Features</a>
        <a href="#gallery">Gallery</a>
        <a href="#contact">Contact</a>
        <!-- External links in nav -->
        <a href="https://www.w3schools.com" target="_blank" rel="noopener noreferrer">W3Schools</a>
        <a href="https://developer.mozilla.org" target="_blank" rel="noopener noreferrer">MDN Web Docs</a>
    </nav>
</header>

<section id="introduction">
    <h2>Introduction</h2>
    <p>This document demonstrates a variety of HTML elements including anchor tags, paragraphs, lists, images, and blockquotes. It is designed to be a comprehensive example for learning purposes.</p>
    <p>Feel free to navigate to different sections using the links in the navigation bar above.</p>
    <p>To learn HTML basics, you might want to visit <a href="https://www.html.com" target="_blank" rel="noopener noreferrer">HTML.com</a>, a great resource for beginners.</p>
</section>

<section id="history">
    <h2>History</h2>
    <p>The use of HTML has evolved significantly since its inception in the early 1990s. Initially designed for simple document markup, it now supports complex applications on the web.</p>
    <p>To learn more about HTML history, you can visit the <a href="https://en.wikipedia.org/wiki/HTML" target="_blank" rel="noopener noreferrer">Wikipedia page on HTML</a>.</p>
    <blockquote>
        "HTML is the backbone of the web. Without it, the internet would be just a collection of disconnected files." – Web Historian
    </blockquote>
    <p>For a timeline of web technology, check out <a href="https://webtimeline.net" target="_blank" rel="noopener noreferrer">WebTimeline.net</a>.</p>
</section>

<section id="features">
    <h2>Features of This Document</h2>
    <ul>
        <li>Anchor tags linking to different sections within the page</li>
        <li>Paragraphs with detailed text</li>
        <li>Images with captions</li>
        <li>Blockquotes for highlighting quotes</li>
        <li>Lists and structured content</li>
    </ul>
    <p>Each feature serves to make the document more user-friendly and accessible.</p>
    <p>Learn more about HTML elements on <a href="https://developer.mozilla.org/en-US/docs/Web/HTML/Element" target="_blank" rel="noopener noreferrer">MDN's HTML Elements Reference</a>.</p>
</section>

<section id="gallery">
    <h2>Gallery</h2>
    <figure>
        <img src="https://via.placeholder.com/600x300.png?text=Sample+Image+1" alt="Sample Image 1" />
        <figcaption>Sample Image 1: A placeholder image representing a sample photo.</figcaption>
    </figure>
    <figure>
        <img src="https://via.placeholder.com/600x300.png?text=Sample+Image+2" alt="Sample Image 2" />
        <figcaption>Sample Image 2: Another placeholder image for demonstration.</figcaption>
    </figure>
    <p>For free images, visit <a href="https://unsplash.com" target="_blank" rel="noopener noreferrer">Unsplash</a> or <a href="https://pixabay.com" target="_blank" rel="noopener noreferrer">Pixabay</a>.</p>
</section>

<section id="contact">
    <h2>Contact Information</h2>
    <p>If you have any questions or would like to provide feedback, please feel free to reach out:</p>
    <address>
        Email: <a href="mailto:info@example.com">info@example.com</a><br />
        Phone: <a href="tel:+1234567890">+1 (234) 567-890</a><br />
        Address: 123 Example St., Sample City, Country
    </address>
    <p>Follow us on social media:</p>
    <ul>
        <li><a href="https://twitter.com/example" target="_blank" rel="noopener noreferrer">Twitter</a></li>
        <li><a href="https://facebook.com/example" target="_blank" rel="noopener noreferrer">Facebook</a></li>
        <li><a href="https://linkedin.com/company/example" target="_blank" rel="noopener noreferrer">LinkedIn</a></li>
    </ul>
</section>

<footer>
    <p>© 2025 Sample HTML Document. All rights reserved.</p>
    <p><a href="#top">Back to top</a></p>
</footer>

</body>
</html>

`

func TestParser(t *testing.T) {
	fmt.Println(parseDocument(document).Text)
}

func BenchmarkParser(b *testing.B) {
	for b.Loop() {
		parseDocument(document)

	}
}
