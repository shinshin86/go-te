package te

func (t *Te) BeforeAll(fn func()) {
	t.beforeAll = fn
}

func (t *Te) AfterAll(fn func()) {
	t.afterAll = fn
}

func (t *Te) BeforeEach(fn func()) {
	t.beforeEach = fn
}

func (t *Te) AfterEach(fn func()) {
	t.afterEach = fn
}
