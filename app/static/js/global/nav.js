(() => {
    const nav = document.querySelector('#nav');
    const mobileNav = document.querySelector('#mobile-nav');
    const breadcrumb = document.querySelector("#breadcrumb");

    const navStateAttr = 'data-open';

    let isNavOpen = false;

    mobileNav.addEventListener('click', () => {
        isNavOpen = !isNavOpen;
        nav.setAttribute(navStateAttr, isNavOpen.toString());
    });

    document.querySelector("body").addEventListener("click", (e) => {
        if (isNavOpen && e.target !== mobileNav && !mobileNav.contains(e.target)) {
            isNavOpen = false;
            nav.setAttribute(navStateAttr, 'false');
        }
    });




    nav.querySelector("ul").querySelectorAll("a").forEach(link => {
        if (link.getAttribute("href") === `/${window.activePage}`) {
            link.setAttribute("data-active", "true");
        }
    });



    const renderBreadcrumb = () => {
        const breadcrumbData = document.querySelector("body").getAttribute("data-breadcrumb");
        const url = window.location.pathname.substring(1).split('/');;

        breadcrumb.innerHTML = '<li><a href="/" title="Home">Home</a></li>' + url.map((item, index) => {
            if (item === "") return;

            if (index === url.length - 1) {
                return `<li>${breadcrumbData}</li>`;
            }

            const name = item.charAt(0).toUpperCase() + item.slice(1);
            return `<li><a href="/${item}" title="${name}">${name}</a></li>`;
        }).join("");
    }

    document.addEventListener('keyup', (e) => {
        if (e.ctrlKey && e.shiftKey && e.key === 'A') {
            e.preventDefault();
            window.location.href = '/admin/';
        }
    }, false);

    document.addEventListener('htmx:afterSwap', function (event) {
        renderBreadcrumb();
    });

    renderBreadcrumb();
})();