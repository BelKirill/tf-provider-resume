Below is an idea for a CV in Terraform

## resource "resume"
An example of a resume resource. This will create the document that will be stored in a document database in a json format.
```
data "intro" "senior" {
  version = "latest"
}

data "summary" "challenge" {
  version = "latest"
}

data "position" "fundbox" {
  aspect = ["devops", "innovation"]
}

data "position" "compugen" {
  aspect = ["devops", "oms"]
}

data "position" "doit" {
  aspect = ["costs", "multicloud"]
}

data "position" "personetics" {
  aspect = ["tech_lead", "devops", "custom"]
}

resource "resume" "github" {
  position_type = "devops"
  status = "alive"
  company = "GitHub"
  
  intro = intro.senior.markdown

  job_history_highlight = [
    position.fundbox.markdown,
    position.compugen.markdown,
    position.doit.markdown,
    position.personetics.markdown,
  ]

  skill_stack = flatten(
    position.fundbox.skills,
    position.compugen.skills,
    position.doit.skills,
    position.personetics.skills,
  )

  summary = summary.challenge.markdown
}
```

## resource "position"
```
resource "position" "fundbox" {
  company = "Fundbox"
  title = "Senior DevOps Engineer"
  dates = "Feb 2022 - Apr 2023"

  description = [
    {
      markdown = ""
      aspects = [
        "",
      ]
    },
  ]
  achievements = [
    {
      markdown = ""
      aspects = [
        "",
      ]
    },
  ]
  skills = [
    {
      skill_name = ""
      aspects = [
        "",
      ]
    },
  ]
}
```

## resource "intro"
```
resource "intro" "senior" {
  markdown = ""
}
```

## resource "summary"
```
resource "summary" "challenge" {
  markdown = ""
}
```
