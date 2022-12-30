# [What is Server Side Rendering - Say hi to Next.js](https://youtu.be/721mxZ3XHf8)

## What is Rendering
Rendering refers to the process of generating a visual representation of a 3D scene or object. This is typically done using computer software, such as a 3D rendering program, to create a 2D image or animation that accurately represents the appearance and visual properties of the scene or object.

Rendering is used in a variety of fields, including architecture, interior design, and film and video game production. It allows designers and artists to create detailed, realistic visualizations of their ideas and concepts, which can be used for presentations, marketing materials, or to visualize the final result of a design project.

There are several different techniques and algorithms that can be used for rendering, each with their own strengths and limitations. Some common rendering techniques include ray tracing, which simulates the way light interacts with objects in a scene; scanline rendering, which generates an image by scanning a 2D grid across the screen; and global illumination, which takes into account the way light bounces off surfaces in a scene to create more realistic lighting and shadows.


## Client Side Rendering
Client-side rendering, also known as client-side rendering, refers to the process of rendering content on the client side of a web application, rather than on the server side.

In a client-side rendering system, a web application is typically written in JavaScript and runs in the user's web browser. When a user requests a web page, the server sends the HTML, CSS, and JavaScript files for the page to the browser. The browser then executes the JavaScript code, which generates the content for the page and updates the HTML DOM (Document Object Model) to reflect the new content.

Client-side rendering has several advantages over server-side rendering. It allows for faster page load times, as the server does not need to generate and send the full HTML for each page request. It also allows for more dynamic and interactive web applications, as the JavaScript code can update the page content in real-time based on user input or other events.

However, client-side rendering can also have some limitations. It can be slower for initial page loads, as the browser must execute the JavaScript code before it can display the content. It can also be less SEO-friendly, as search engines may have difficulty indexing dynamically generated content.


## Pre-rendering
Pre-rendering, also known as server-side rendering, is a technique for rendering content on the server side of a web application before it is sent to the client's web browser.

In a pre-rendering system, the server generates the full HTML for a web page and sends it to the client's browser, along with any necessary CSS and JavaScript files. The browser can then display the page immediately, without having to execute any additional JavaScript code.

Pre-rendering can be used to improve the performance and user experience of a web application, particularly for applications that rely heavily on client-side rendering. It allows the page to load faster, as the browser does not have to wait for the JavaScript code to execute before it can display the content. Pre-rendering can also improve SEO, as search engines can easily index the static HTML content of the page.

However, pre-rendering can also have some limitations. It can be more resource-intensive for the server, as it must generate the full HTML for each page request. It can also make it more difficult to create dynamic and interactive web applications, as the content must be generated on the server side rather than in the browser.


## Static Site Generation
Static site generation is a process for building websites in which the content is generated as a set of static HTML, CSS, and JavaScript files, rather than being generated dynamically on the server side.

In a static site generation system, the content for a website is typically stored in a content management system (CMS) or a flat file database, and a build process is used to generate the static HTML, CSS, and JavaScript files for the site. The static files are then uploaded to a web server, where they can be served to users.

Static site generation has several advantages over traditional dynamic websites, which rely on server-side rendering to generate the HTML for each page request. Static sites are typically faster and more lightweight, as they do not require a server to generate the content on each page request. They are also easier to scale and more secure, as there is no server-side code that can be hacked or compromised.

However, static site generation can also have some limitations. It can be more difficult to create dynamic and interactive websites, as the content is static and cannot be updated in real-time. It can also require more setup and maintenance, as the build process must be run whenever the content is updated.


## Server Side Rendering
Server-side rendering (SSR) is a technique for rendering content on the server side of a web application, rather than on the client side.

In a server-side rendering system, a web application is typically written in a server-side language, such as PHP or Ruby, and runs on a web server. When a user requests a web page, the server generates the HTML for the page and sends it to the user's web browser. The browser then displays the page, along with any necessary CSS and JavaScript files.

Server-side rendering has several advantages over client-side rendering, which generates the content on the client side using JavaScript. It allows for faster initial page load times, as the server generates the full HTML for the page and sends it to the browser, rather than waiting for the JavaScript code to execute. It can also be more SEO-friendly, as search engines can easily index the static HTML content of the page.

However, server-side rendering can also have some limitations. It can be more resource-intensive for the server, as it must generate the HTML for each page request. It can also be less dynamic and interactive, as the content must be generated on the server side rather than in the browser.


## Primitive Wrapper Type
In some programming languages, a primitive wrapper type is a data type that wraps or encapsulates a primitive data type, such as an integer or a boolean value. Primitive wrapper types are often used to provide additional functionality or to allow primitive values to be used in contexts where only objects are allowed.

For example, in the Java programming language, the primitive types int, boolean, and char are wrapped by the Integer, Boolean, and Character wrapper classes, respectively. These wrapper classes provide methods for converting primitive values to and from strings, as well as other utility methods.

In languages that support object-oriented programming, primitive wrapper types can be useful for storing primitive values in collections or for passing primitive values as method arguments when objects are expected. However, they can also add overhead and complexity, as they require the creation of additional objects and can result in increased memory usage.


## Introduction to Next.js
Next.js is a popular open-source JavaScript framework for building server-rendered, statically generated, and client-side rendered web applications. It is built on top of React, a popular JavaScript library for building user interfaces, and is designed to make it easy to build scalable, performance-optimized web applications.

One of the key features of Next.js is server-side rendering, which allows the server to generate the initial HTML for a page and send it to the client's browser. This can improve the performance and user experience of a web application, particularly for applications with a large amount of content or for users with slow internet connections.

Next.js also includes a powerful static site generation feature, which allows developers to build and deploy their applications as a set of static HTML, CSS, and JavaScript files. This can improve the performance and scalability of a web application, as it reduces the load on the server and allows the application to be served from a content delivery network (CDN).

In addition to server-side rendering and static site generation, Next.js also provides a range of other features and tools for building web applications, including automatic code splitting, automatic optimization of images and other assets, and support for serverless functions.


## How Next.js works
Next.js is a framework for building web applications with server-side rendering, statically generated pages, and client-side rendering. It works by combining the power of React, a popular JavaScript library for building user interfaces, with a set of tools and conventions for building web applications.

Here's a general overview of how Next.js works:

- The developer creates a Next.js project and writes the application code using React components.

- The Next.js build process generates a set of static HTML, CSS, and JavaScript files for the application, based on the code and configuration of the project.

- The static files are uploaded to a web server or deployed to a content delivery network (CDN).

- When a user requests a page from the application, the server or CDN responds with the static HTML, CSS, and JavaScript files for the page.

- The user's web browser loads and executes the JavaScript code, which renders the React components and updates the page content.


In a Next.js application, the server-side rendering and static generation features are optional and can be enabled or disabled based on the needs of the application. This allows developers to choose the most appropriate rendering approach for their application, depending on the requirements for performance, scalability, and interactivity.


## SPA in Next.js
A single-page application (SPA) is a type of web application that loads a single HTML page and dynamically updates the content of the page as the user interacts with the application. This can provide a more seamless and responsive user experience, as the application does not need to reload the page or navigate to new pages as the user moves between different sections of the application.

Next.js is a framework that can be used to build single-page applications (SPAs) using React and server-side rendering. In a Next.js SPA, the server generates the initial HTML for the page and sends it to the client's browser, along with the necessary CSS and JavaScript files. The JavaScript code then takes over, rendering the React components and updating the page content as the user interacts with the application.

To build an SPA with Next.js, the developer can create a Next.js project and write the application code using React components. The Next.js build process will then generate the static HTML, CSS, and JavaScript files for the application, which can be deployed to a web server or CDN. When a user requests a page from the application, the server or CDN will respond with the static files, and the JavaScript code will execute in the user's browser to render the React components and update the page content.

## Next.js vs Gatsby
Next.js and Gatsby are both popular open-source JavaScript frameworks for building web applications. Both frameworks are built on top of React, a popular library for building user interfaces, and provide a range of tools and features for building performance-optimized web applications.

However, Next.js and Gatsby have some key differences in terms of their approach to building web applications:

- Server-side rendering vs. static site generation: Both Next.js and Gatsby support server-side rendering, but they differ in how they handle static site generation. Next.js provides both server-side rendering and static site generation as optional features, allowing developers to choose the most appropriate rendering approach for their application based on the requirements for performance, scalability, and interactivity. Gatsby, on the other hand, is primarily focused on static site generation, using GraphQL and a build process to generate a set of static HTML, CSS, and JavaScript files for the application.


- Deployment options: Next.js applications can be deployed to a variety of hosting platforms, including traditional web servers, cloud platforms, and serverless environments. Gatsby applications, on the other hand, are typically deployed as a set of static files, which can be served from a variety of static file hosting platforms or integrated with a traditional web server.

- Ecosystem and community: Both Next.js and Gatsby have active communities and ecosystems of plugins, integrations, and templates, but the focus and scope of these communities may differ. Next.js is part of the larger React ecosystem and is often used for building traditional web applications with server-side rendering and client-side rendering. Gatsby, on the other hand, is often used for building static site generators and Jamstack applications, and has a strong focus on performance, security, and integrations with modern web technologies.

## Next.js Features
Next.js is a popular open-source JavaScript framework for building web applications with server-side rendering, statically generated pages, and client-side rendering. It is built on top of React and provides a range of features and tools to make it easier to build performance-optimized web applications.

Some of the key features of Next.js include:

__Feature of Next js__
- Scalable
- Performance
- Developer experience
- Webpack
- Fast Refresh
- Enhanced Debugging
- Typescript
- Static Assets Ready
- Image Optimization
- Internationalization
- Redirect / Rewrites
- Deployment

- Server-side rendering: Next.js allows developers to build web applications with server-side rendering, which allows the server to generate the initial HTML for a page and send it to the client's browser. This can improve the performance and user experience of a web application, particularly for applications with a large amount of content or for users with slow internet connections.

- Static site generation: Next.js includes a powerful static site generation feature, which allows developers to build and deploy their applications as a set of static HTML, CSS, and JavaScript files. This can improve the performance and scalability of a web application, as it reduces the load on the server and allows the application to be served from a content delivery network (CDN).

- Automatic code splitting: Next.js automatically splits the code for a web application into smaller chunks, based on the routes and components used by the application. This can improve the performance of the application, as the browser only needs to download and execute the code required for the current page, rather than loading the entire application at once.

- Automatic optimization of assets: Next.js automatically optimizes images and other assets used in a web application, including resizing and compressing images and converting font files to web formats. This can improve the performance and loading times of the application.

- Serverless functions: Next.js supports the use of serverless functions, which allow developers to run code on demand in response to specific events, such as a user request or a change to a database. This can be used to build scalable and cost-effective web applications.