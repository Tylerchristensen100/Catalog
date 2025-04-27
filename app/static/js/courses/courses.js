(function () {
    const aside = document.querySelector('aside');
    const filtersContainer = document.querySelector('#course-filters');
    const filterAccordion = document.querySelector('#filter-accordion');

    const search = document.querySelector('#search input');
    const autoComplete = document.querySelector('#search #autocomplete');


    const uniqueMajors = new Set();
    let majorCodes = [];
    filtersContainer.innerHTML = "<h2>Filters</h2>";
    window.courses.forEach(course => {
        if (course.major && !uniqueMajors.has(course.major.code)) {
            if (course.major.code.length > 0 && course.major.id > 0) {
                uniqueMajors.add(course.major.code);
                majorCodes.push(course.major);
            }
        }
    });
    majorCodes = Array.from(majorCodes).sort();


    filtersContainer.innerHTML = majorCodes.map(majorCode => {
        return `<div>
                    <input type="checkbox" id="${majorCode.id}" name="${majorCode.name}" value="${majorCode.code}" checked="true">
                    <label for="${majorCode.id}">${majorCode.name}</label>
                </div>`
    }).join('');

    const filters = document.querySelectorAll('aside input[type="checkbox"]');

    filters.forEach((filter) => {
        filter.addEventListener('change', () => {
            const checkedFilters = Array.from(filters).filter(f => f.checked).map(f => f.value);
            const courses = document.querySelectorAll('.course');

            courses.forEach(course => {
                const courseMajorCode = course.getAttribute('data-major-code');
                if (checkedFilters.length === filters.length || checkedFilters.includes(courseMajorCode)) {
                    course.style.display = 'block';
                } else {
                    course.style.display = 'none';
                }
            });

            document.querySelectorAll('h3[data-major-code]').forEach((h3) => {
                const courseMajorCode = h3.getAttribute('data-major-code');
                if (checkedFilters.length === filters.length || checkedFilters.includes(courseMajorCode)) {
                    h3.style.display = 'block';
                } else {
                    h3.style.display = 'none';
                }
            });
        });
    });







    const filterCourses = () => {
        const searchTerm = search.value.toLowerCase();

        const filteredCourses = searchTerm.length > 0 ?
            window.courses.filter((course) => {
                const name = course.name.toLowerCase();
                const majorCode = course.major_code.toLowerCase();
                const code = course.code.toLowerCase();
                return name.includes(searchTerm) || majorCode.includes(searchTerm) || code.includes(searchTerm);
            }) : window.courses;



        const filteredList = filteredCourses.map((c) => {
            return `<li key="${c.id}" class="course" data-major-code="${c.major_code}" data-code="${c.code}"
                data-credit-hours="${c.creditHours}">
                <a href="/courses/${c.major_code}-${c.code}" title="${c.name}">
                    ${c.major_code} - ${c.code} | ${c.name}
                </a>
            </li>`;
        }).join('');

        if (filteredList.length > 0) {
            autoComplete.style.display = 'block';
            autoComplete.innerHTML = `<ul>${filteredList}</ul> `;
        } else {
            autoComplete.innerHTML = '<li>No results found</li>';
        }
    }



    search.addEventListener('keyup', (e) => {
        filterCourses()
    });

    search.addEventListener('focus', () => {
        autoComplete.style.display = 'block';
        filterCourses();
    })

    search.addEventListener('blur', () => {
        setTimeout(() => {
            search.value = '';
            autoComplete.style.display = 'none';
        }, 200);
    });

    filterAccordion.addEventListener('click', (e) => {
        const isOpen = aside.getAttribute('data-open') === 'true';
        aside.setAttribute('data-open', !isOpen);
        aside.setAttribute('aria-expanded', !isOpen);
    });


})()