(function () {
    const filtersContainer = document.querySelector("aside");
    const filterAccordion = document.querySelector("#filter-accordion");
  
    const degreeFilters = document.querySelectorAll(
      'aside #degree input[type="checkbox"]'
    );
    const schoolFilters = document.querySelectorAll(
      'aside #school input[type="checkbox"]'
    );
    const search = document.querySelector("#search input");
    const autoComplete = document.querySelector("#search #autocomplete");
  
    const allFilters = [...degreeFilters, ...schoolFilters];
  
    const filterPrograms = () => {
      const checkedDegrees = Array.from(degreeFilters)
        .filter((f) => f.checked)
        .map((f) => f.value);
      const checkedSchools = Array.from(schoolFilters)
        .filter((f) => f.checked)
        .map((f) => f.value);
      const programs = document.querySelectorAll(".program");
  
      programs.forEach((program) => {
        const programDegree = program.getAttribute("data-grad-level");
        const programSchool = program.getAttribute("data-school");
  
        const degreesMatch =
          checkedDegrees.length === degreeFilters.length ||
          checkedDegrees.includes(programDegree);
        const schoolsMatch =
          checkedSchools.length === schoolFilters.length ||
          checkedSchools.includes(programSchool.toLowerCase());
  
        if (degreesMatch && schoolsMatch) {
          program.style.display = "block";
        } else {
          program.style.display = "none";
        }
      });
    };
  
    allFilters.forEach((f) => f.addEventListener("change", filterPrograms));
  
    const searchPrograms = () => {
      const searchTerm = search.value.toLowerCase();
  
      const filteredPrograms =
        searchTerm.length > 0
          ? window.programs.filter((program) => {
              const name = program.name.toLowerCase();
              const description = program.description.toLowerCase();
              return (
                name.includes(searchTerm) || description.includes(searchTerm)
              );
            })
          : window.programs;
  
      const filteredList = filteredPrograms
        .map((p) => {
          return `<li key="${p.id}" class="program" data-major-code="${
            p.major_code
          }" data-grad-level="${p.grad_level}"
                  data-program-type="${p.program_type}" data-school="${
            p.school
          }" data-online="${p.online > 0}">
                  <a href="/programs/${p.name}" title="${p.name}">
                  ${p.name}
                  </a>
              </li>`;
        })
        .join("");
  
      if (filteredList.length > 0) {
        autoComplete.style.display = "block";
        autoComplete.innerHTML = `<ul>${filteredList}</ul>`;
      } else {
        autoComplete.innerHTML = "<li>No results found</li>";
      }
    };
  
    search.addEventListener("keyup", (e) => {
      searchPrograms();
    });
  
    search.addEventListener("focus", () => {
      autoComplete.style.display = "block";
      searchPrograms();
    });
  
    search.addEventListener("blur", () => {
      setTimeout(() => {
        search.value = "";
        autoComplete.style.display = "none";
      }, 200);
    });
  
    filterAccordion.addEventListener("click", (e) => {
      const isOpen = filtersContainer.getAttribute("data-open") === "true";
      filtersContainer.setAttribute("data-open", !isOpen);
      filterAccordion.setAttribute("aria-expanded", !isOpen);
    });
  })();