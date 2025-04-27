import styles from "./footer.module.css";

function Footer() {
  return (
    <footer className={styles.footer}>
      <nav>
        <ul className="">
        <li>
            <a href="/">Public Site</a>
          </li>
          <li>
            <a href="/about">About</a>
          </li>
          <li>
            <a href="/contact">Contact</a>
          </li>
          <li>
            <a href="/privacy">Privacy Policy</a>
          </li>
        </ul>
      </nav>
      <div className={styles.copyright}>
        <p>
          &copy; {new Date().getFullYear()} Tyler Christensen. All rights
          reserved.
        </p>
      </div>
    </footer>
  );
}

export default Footer;
