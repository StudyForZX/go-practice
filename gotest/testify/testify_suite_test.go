package testify

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
	indent int
}

func (suite *ExampleSuite) indents() (result string) {
	for i := 0; i < suite.indent; i++ {
		result += "----"
	}
	return
}

// 首先执行
// SetupAllSuite has a SetupSuite method, which will run before the tests in the suite are run.
// 该方法在整个测试套件中的所有测试开始之前只执行一次。
// 通常用于执行整个测试套件只需设置一次的初始化操作，比如数据库连接、昂贵的资源初始化等。
func (suite *ExampleSuite) SetupSuite() {
	fmt.Println("Suite setup")
}

// 最后执行
func (suite *ExampleSuite) TearDownSuite() {
	fmt.Println("Suite teardown")
}

// 第二个执行
// SetupTestSuite has a SetupTest method, which will run before each test in the suite.
// 该方法在测试套件中的每个测试开始前都会执行。
// 通常用于每个测试都需要重新设置的操作，比如测试数据的准备、测试环境的清理等。
func (suite *ExampleSuite) SetupTest() {
	suite.indent++
	fmt.Println(suite.indents(), "Test setup")
}

// 倒数第二执行
func (suite *ExampleSuite) TearDownTest() {
	fmt.Println(suite.indents(), "Test teardown")
	suite.indent--
}

// 第三个执行
func (suite *ExampleSuite) BeforeTest(suiteName, testName string) {
	suite.indent++
	fmt.Printf("%sBefore %s.%s\n", suite.indents(), suiteName, testName)
}

// 倒数第三执行
func (suite *ExampleSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("%sAfter %s.%s\n", suite.indents(), suiteName, testName)
	suite.indent--
}

// SetupSubTest has a SetupSubTest method, which will run before each subtest in the suite.
func (suite *ExampleSuite) SetupSubTest() {
	suite.indent++
	fmt.Println(suite.indents(), "SubTest setup")
}

func (suite *ExampleSuite) TearDownSubTest() {
	fmt.Println(suite.indents(), "SubTest teardown")
	suite.indent--
}

func (suite *ExampleSuite) TestCase1() {

	suite.indent++

	defer func() {
		fmt.Println(suite.indents(), "End TestCase1")
		suite.indent--
	}()

	fmt.Println(suite.indents(), "Begin TestCase1")

	suite.Run("case1-subtest1", func() {
		suite.indent++
		fmt.Println(suite.indents(), "Begin TestCase1.Subtest1")
		fmt.Println(suite.indents(), "End TestCase1.Subtest1")
		suite.indent--
	})

	suite.Run("case2-subtest2", func() {
		suite.indent++
		fmt.Println(suite.indents(), "Begin TestCase1.Subtest2")
		fmt.Println(suite.indents(), "End TestCase1.Subtest2")
		suite.indent--
	})
}

func (suite *ExampleSuite) TestCase2() {

	suite.indent++

	defer func() {
		fmt.Println(suite.indents(), "End TestCase2")
		suite.indent--
	}()

	fmt.Println(suite.indents(), "Begin TestCase2")

	suite.Run("case2-subtest1", func() {
		suite.indent++
		fmt.Println(suite.indents(), "Begin TestCase2.SubTest1")
		fmt.Println(suite.indents(), "End TestCase2.SubTest1")
		suite.indent--
	})
}

func TestExampleSuite(t *testing.T) {
	// suite.Run函数执行的大致逻辑是：通过反射机制得到了目标类型的方法集合，并执行集合中所有前缀为Test开头的方法。
	suite.Run(t, new(ExampleSuite))
}

// OutPut is below

// Suite setup
// ---- Test setup
// --------Before ExampleSuite.TestCase1
// ------------ Begin TestCase1
// ---------------- SubTest setup
// -------------------- Begin TestCase1.Subtest1
// -------------------- End TestCase1.Subtest1
// ---------------- SubTest teardown
// ---------------- SubTest setup
// -------------------- Begin TestCase1.Subtest2
// -------------------- End TestCase1.Subtest2
// ---------------- SubTest teardown
// ------------ End TestCase1
// --------After ExampleSuite.TestCase1
// ---- Test teardown
// ---- Test setup
// --------Before ExampleSuite.TestCase2
// ------------ Begin TestCase2
// ---------------- SubTest setup
// -------------------- Begin TestCase2.SubTest1
// -------------------- End TestCase2.SubTest1
// ---------------- SubTest teardown
// ------------ End TestCase2
// --------After ExampleSuite.TestCase2
// ---- Test teardown
// Suite teardown

// 下边不是很懂

// 我们知道：
// go test工具可以通过-run命令行参数来选择要执行的TestXXX函数，考虑到testify使用TestXXX函数拉起测试套件(XXXSuite)，
// 因此从testify视角来看，通过go test -run可以选择执行哪个XXXSuite，
// 前提是一个TestXXX中仅初始化和运行一种XXXSuite的所有测试用例。

// 如果要选择XXXSuite的方法(即testify眼中的测试用例)，我们不能用-run了，
// 需要使用testify新增的-m命令行选项，下面是一个仅执行带有Case2关键字测试用例的示例：

// go test -testify.m Case2

// Suite setup
// ---- Test setup
// --------Before ExampleSuite.TestCase2
// ------------ Begin TestCase2
// ---------------- SubTest setup
// -------------------- Begin TestCase2.SubTest1
// -------------------- End TestCase2.SubTest1
// ---------------- SubTest teardown
// ------------ End TestCase2
// --------After ExampleSuite.TestCase2
// ---- Test teardown
// Suite teardown
// PASS
// ok      github.com/studyforzx/lt/gotest 1.945s
