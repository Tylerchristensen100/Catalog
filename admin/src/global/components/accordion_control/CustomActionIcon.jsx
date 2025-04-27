import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';


const CustomActionButton = ({ icon, onClick, className }) => {
    return (
        <div
            onClick={onClick}
            className={`custom_action_button ${className}`}
        >
            <FontAwesomeIcon icon={icon} />
        </div>
    );
};

export default CustomActionButton;