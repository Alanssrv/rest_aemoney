CREATE TABLE `Transaction`(
    TransactionId INT AUTO_INCREMENT,
    `Name` VARCHAR(255),
    `Value` DECIMAL,
    `Type` TINYINT,
    Category VARCHAR(255),
    CreationDate DATE,
    UpdateDate DATE,
    DeletionDate DATE,
    CONSTRAINT PK_Transaction PRIMARY KEY (TransactionId)
);

CREATE PROCEDURE GetAllTransactions()
BEGIN
    SELECT * from `Transaction` WHERE DeletionDate IS NULL;
END;

CREATE PROCEDURE GetTransactionById(paramTransactionId INT)
BEGIN
    SELECT * FROM `Transaction` WHERE DeletionDate IS NULL AND TransactionId = paramTransactionId;
END;

CREATE PROCEDURE InsertTransaction(`Name` VARCHAR(255), `Value` DECIMAL, `Type` TINYINT, Category VARCHAR(255))
BEGIN
    INSERT INTO `Transaction` (Name, Value, Type, Category, CreationDate, UpdateDate, DeletionDate)
    VALUES (
        `Name`,
        `Value`,
        `Type`,
        Category,
        NOW(),
        NOW(),
        NULL
    );
END;

CREATE PROCEDURE UpdateTransaction(`newName` VARCHAR(255), `newValue` DECIMAL, `newType` TINYINT, newCategory VARCHAR(255), paramTransactionId INT)
BEGIN
    UPDATE `Transaction` 
    SET
        `Name` =  `newName`,
        `Value` = `newValue`,
        `Type` = `newType`,
        Category = newCategory,
        UpdateDate = NOW()
    WHERE DeletionDate IS NULL AND TransactionId = paramTransactionId;
END;

CREATE PROCEDURE DeleteTransactionByID(paramTransactionId INT)
BEGIN
    UPDATE `Transaction`
    SET DeletionDate = NOW()
    WHERE DeletionDate IS NULL AND TransactionId = paramTransactionId;
END;