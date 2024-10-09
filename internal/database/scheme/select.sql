SELECT card.id, card.person_inn, person.name, card.bank_bic, bank.name, card.valid
FROM card
INNER JOIN person ON card.person_inn = person.inn
INNER JOIN bank ON card.bank_bic = bank.bic
WHERE card.id=111 AND person_inn = 123 AND bank.bic = 456 AND card.valid = '2020-01-01'