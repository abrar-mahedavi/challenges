`Medium`	`Codewriting` 	`300`

Your company is planning to expand internationally very soon. You have been tasked with preparing a report on foreign markets and potential competitors.

After some investigation, you've created a database containing a foreignCompetitors table, which has the following structure:

- competitor: the name of the competitor;
- country: the country in which the competitor is operating.

In your report, you need to include the number of competitors per country and an additional row at the bottom that contains a summary: ("Total:", total_number_of_competitors)

Given the foreignCompetitors table, compose the resulting table with two columns: country and competitors. The first column should contain the country name, and the second column should contain the number of competitors in this country. The table should be sorted by the country names in ascending order. In addition, it should have an extra row at the bottom with the summary, as described above.

## Example

- For the following table soccer_team

| id   | first_name | surname  | player_number |
|------|------------|----------|---------------| 
| 1    | Alexis     | Sanchez  | 7             | 
| 2    | Petr       | Cech     | 33            | 
| 3    | Hector     | Bellerin | 24            | 
| 4    | Olivier    | Giroud   | 12            | 
| 5    | Theo       | Walcott  | 14            | 
| 6    | Santi      | Cazorla  | 19            | 

- the output should be

| players                                                                                                       |
|---------------------------------------------------------------------------------------------------------------|
| Alexis Sanchez #7; Oliver Giroud #12; Theo Walcott #14; Santi Cazorla #19; Hector Bellerin #24; Petr Cech #33 |

# Output
- [execution time limit] 10 seconds (mysql)

