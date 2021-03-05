package kreta

func (s *Session) Grades() (*Grades, error) {}

func (s *Session) Absences() (*Absences, error) {}

func (s *Session) Timetable() (*Timetable, error) {}

func (s *Session) Notes() (*Notes, error) {}

func (s *Session) Notices() (*Notices, error) {}

func (s *Session) Tests() (*Tests, error) {}

func (s *Session) Groups() (*Groups, error) {}

func (s *Session) UserData() (*UserData, error) {}

func (s *Session) YearOrder() (*YearOrder, error) {}

func (s *Session) Homework() (*Homework, error) {}

func (s *Session) HomeworkComments(id string) (*HomeworkComments, error) {}

func (s *Session) HomeworkComment(id, comment string) error {}
