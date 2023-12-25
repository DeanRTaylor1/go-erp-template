interface TablePaginationProps {
    page: number; // current page number
    count: number; // total number of rows
    rowsPerPage: number; // number of rows per page
    onPageChange: (newPage: number) => void;
    onRowsPerPageChange: (newRowsPerPage: number) => void;
}

const TablePagination: React.FC<TablePaginationProps> = ({ page, count, rowsPerPage, onPageChange, onRowsPerPageChange }) => {
    const totalPages = Math.ceil(count / rowsPerPage);

    const handlePreviousPage = () => {
        if (page > 0) onPageChange(page - 1);
    };

    const handleNextPage = () => {
        if (page < totalPages - 1) onPageChange(page + 1);
    };

    return (
        <div className="w-full flex gap-4 pt-2">

            <div className="ml-auto flex gap-4 items-center">
                <span>{`${page + 1}-${count <= (page + 1) * rowsPerPage ? count : (page + 1) * rowsPerPage} of ${count}`}</span>
                <div>
                    <label>
                        Rows per page:
                        <select
                            value={rowsPerPage}
                            onChange={(e) => onRowsPerPageChange(Number(e.target.value))}
                            className="ml-1"
                        >
                            <option value={5}>5</option>
                            <option value={10}>10</option>
                            <option value={25}>25</option>
                        </select>
                    </label>
                </div>
                <button
                    onClick={handlePreviousPage}
                    disabled={page === 0}
                    className="p-1"
                >
                    &lt;
                </button>
                <button
                    onClick={handleNextPage}
                    disabled={page >= totalPages - 1}
                    className="p-1"
                >
                    &gt;
                </button>
            </div>
        </div>
    );

};

export default TablePagination;
